package manager

import (
	"context"
	"sync"
	"time"

	"github.com/RandalTeng/go-oauth2-server/definition"
	"github.com/RandalTeng/go-oauth2-server/errors"
	"github.com/RandalTeng/go-oauth2-server/generator"
)

// NewDefaultManager create to default authorization management instance
func NewDefaultManager() *Manager {
	m := NewManager()
	// default implementation
	m.MapAuthorizeGenerate(generator.NewAuthorizeGenerate())
	m.MapAccessGenerate(generator.NewAccessGenerate())

	return m
}

// NewManager create to authorization management instance
func NewManager() *Manager {
	return &Manager{
		gtcfg:       make(map[definition.GrantType]*Config),
		gtCfgLock:   &sync.RWMutex{},
		validateURI: DefaultValidateURI,
	}
}

// Manager provide authorization management
type Manager struct {
	codeExp           time.Duration
	gtcfg             map[definition.GrantType]*Config
	gtCfgLock         *sync.RWMutex // this lock is unnecessary, but lock it anyway, for some illegal usage.
	rcfg              *RefreshingConfig
	validateURI       ValidateURIHandler
	authorizeGenerate definition.AuthorizeGenerate
	accessGenerate    definition.AccessGenerate
	tokenStore        definition.TokenStore
	clientStore       definition.ClientStore
	logger            definition.Logger

	code  definition.CodeInfoAdapter
	token definition.TokenInfoAdapter
}

// get grant type config
func (m *Manager) grantConfig(gt definition.GrantType) *Config {
	m.gtCfgLock.RLock()
	defer m.gtCfgLock.RUnlock()
	if c, ok := m.gtcfg[gt]; ok && c != nil {
		return c
	}
	switch gt {
	case definition.AuthorizationCode:
		return DefaultAuthorizeCodeTokenCfg
	case definition.Implicit:
		return DefaultImplicitTokenCfg
	case definition.PasswordCredentials:
		return DefaultPasswordTokenCfg
	case definition.ClientCredentials:
		return DefaultClientTokenCfg
	}
	return &Config{}
}

// SetAuthorizeCodeExp set the authorization code expiration time
func (m *Manager) SetAuthorizeCodeExp(exp time.Duration) {
	m.codeExp = exp
}

// SetTokenCfgWithType set the authorization code grant token config
func (m *Manager) SetTokenCfgWithType(gt definition.GrantType, cfg *Config) {
	m.gtCfgLock.Lock()
	defer m.gtCfgLock.Unlock()
	m.gtcfg[gt] = cfg
}

// SetRefreshTokenCfg set the refreshing token config
func (m *Manager) SetRefreshTokenCfg(cfg *RefreshingConfig) {
	m.rcfg = cfg
}

// SetValidateURIHandler set the validates that RedirectURI is contained in baseURI
func (m *Manager) SetValidateURIHandler(handler ValidateURIHandler) {
	m.validateURI = handler
}

// SetLogger set manager's logger.
func (m *Manager) SetLogger(logger definition.Logger) {
	m.logger = logger
}

// log message with level if Manager.logger is not nil
func (m *Manager) log(level definition.LogLevel, format string, args ...any) {
	if m.logger == nil {
		return
	}
	var logger func(string, ...any)
	switch level {
	case definition.LogLevelFalat:
		logger = m.logger.Falatf
	case definition.LogLevelError:
		logger = m.logger.Errorf
	case definition.LogLevelInfo:
		logger = m.logger.Infof
	case definition.LogLevelTrace:
		if tracer, ok := m.logger.(definition.TraceableLogger); !ok {
			return
		} else {
			logger = tracer.Tracef
		}
	}
	logger(format, args...)
}

// MapAuthorizeGenerate mapping the authorize code generate interface
func (m *Manager) MapAuthorizeGenerate(gen definition.AuthorizeGenerate) {
	m.authorizeGenerate = gen
}

// MapAccessGenerate mapping the access token generate interface
func (m *Manager) MapAccessGenerate(gen definition.AccessGenerate) {
	m.accessGenerate = gen
}

// MapClientStorage mapping the client store interface
func (m *Manager) MapClientStorage(stor definition.ClientStore) {
	m.clientStore = stor
}

// MustClientStorage mandatory mapping the client store interface
func (m *Manager) MustClientStorage(stor definition.ClientStore, err error) {
	if err != nil {
		panic(err.Error())
	}
	m.clientStore = stor
}

// MapTokenStorage mapping the token store interface
func (m *Manager) MapTokenStorage(stor definition.TokenStore) {
	m.tokenStore = stor
}

// MustTokenStorage mandatory mapping the token store interface
func (m *Manager) MustTokenStorage(stor definition.TokenStore, err error) {
	if err != nil {
		panic(err)
	}
	m.tokenStore = stor
}

// GetClient get the client information
func (m *Manager) GetClient(ctx context.Context, clientID string) (cli definition.ClientInfo, err error) {
	cli, err = m.clientStore.GetByID(ctx, clientID)
	if err != nil {
		return
	} else if cli == nil {
		err = errors.ErrInvalidClient
	}
	return
}

// SetCodeAdapter auth code model adapter
func (m *Manager) SetCodeAdapter(adp definition.CodeInfoAdapter) {
	m.code = adp
}

// SetTokenAdapter access token model adapter
func (m *Manager) SetTokenAdapter(adp definition.TokenInfoAdapter) {
	m.token = adp
}

// GenerateAuthToken generate the authorization token(code)
func (m *Manager) GenerateAuthToken(ctx context.Context, rt definition.ResponseType, tgr *definition.TokenGenerateRequest) (any, error) {
	cli, err := m.GetClient(ctx, tgr.ClientID)
	if err != nil {
		return nil, err
	} else if tgr.RedirectURI != "" {
		if err := m.validateURI(cli.GetDomain(), tgr.RedirectURI); err != nil {
			return nil, err
		}
	}

	var info any
	createAt := time.Now()
	td := &definition.GenerateBasic{
		Client:    cli,
		UserID:    tgr.UserID,
		Scope:     tgr.Scope,
		CreatedAt: createAt,
		Request:   tgr.Request,
	}
	switch rt {
	case definition.Code:
		// first step, generate auth code.
		ci := m.code.New()
		ci.SetClientID(tgr.ClientID)
		ci.SetUserID(tgr.UserID)
		ci.SetRedirectURI(tgr.RedirectURI)
		ci.SetScope(tgr.Scope)
		ci.SetExpiredAt(createAt.Add(DefaultCodeExp))
		td.ExpiredAt = ci.GetExpiredAt()
		if tgr.CodeChallenge != "" {
			ci.SetChallenge(tgr.CodeChallenge)
			ci.SetChallengeMethod(tgr.CodeChallengeMethod)
		}

		tv, err := m.authorizeGenerate.Token(ctx, td)
		if err != nil {
			return nil, err
		}
		ci.SetCode(tv)
		info = ci
	case definition.Token:
		// second step, generate access token.
		ti := m.token.New()
		ti.SetClientID(tgr.ClientID)
		ti.SetUserID(tgr.UserID)
		icfg := m.grantConfig(definition.Implicit)
		ti.SetExpiredAt(createAt.Add(icfg.AccessTokenExp))
		td.ExpiredAt = ti.GetExpiredAt()

		body, err := m.accessGenerate.Token(ctx, td, icfg.IsGenerateRefresh)
		if err != nil {
			return nil, err
		}
		ti.SetAccess(body.Access)
		ti.SetAccessIdentifier(body.AccessIdentifier)
		if icfg.IsGenerateRefresh && body.Refresh != "" {
			ti.SetRefreshExpiredAt(createAt.Add(icfg.RefreshTokenExp))
			ti.SetRefresh(body.Refresh)
		}
		info = ti
	}

	err = m.tokenStore.Create(ctx, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (m *Manager) getAuthorizationCode(ctx context.Context, code string) (definition.CodeInfo, error) {
	ci, err := m.tokenStore.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	} else if ci == nil || ci.GetCode() != code ||
		ci.GetExpiredAt().Before(time.Now()) {
		return nil, errors.ErrInvalidAuthorizeCode
	}
	return ci, nil
}

func (m *Manager) delAuthorizationCode(ctx context.Context, code string) error {
	return m.tokenStore.RemoveByCode(ctx, code)
}

func (m *Manager) getAndDelAuthorizationCode(ctx context.Context, tgr *definition.TokenGenerateRequest) (definition.CodeInfo, error) {
	code := tgr.Code
	ci, err := m.getAuthorizationCode(ctx, code)
	if err != nil {
		return nil, err
	} else if ci.GetClientID() != tgr.ClientID {
		return nil, errors.ErrInvalidAuthorizeCode
	} else if codeURI := ci.GetRedirectURI(); codeURI != "" && codeURI != tgr.RedirectURI {
		return nil, errors.ErrInvalidAuthorizeCode
	}

	err = m.delAuthorizationCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return ci, nil
}

func (m *Manager) validateCodeChallenge(ci definition.CodeInfo, ver string) error {
	cc := ci.GetChallenge()
	// early return
	if cc == "" && ver == "" {
		return nil
	}
	if cc == "" {
		return errors.ErrMissingCodeVerifier
	}
	if ver == "" {
		return errors.ErrMissingCodeVerifier
	}
	ccm := ci.GetChallengeMethod()
	if ccm.String() == "" {
		ccm = definition.CodeChallengePlain
	}
	if !ccm.Validate(cc, ver) {
		return errors.ErrInvalidCodeChallenge
	}
	return nil
}

// GenerateAccessToken generate the access token
func (m *Manager) GenerateAccessToken(ctx context.Context, gt definition.GrantType, tgr *definition.TokenGenerateRequest) (definition.TokenInfo, error) {
	cli, err := m.GetClient(ctx, tgr.ClientID)
	if err != nil {
		return nil, err
	}
	if cliPass, ok := cli.(definition.ClientPasswordVerifier); ok {
		if !cliPass.VerifyPassword(tgr.ClientSecret) {
			return nil, errors.ErrInvalidClient
		}
	} else if len(cli.GetSecret()) > 0 && tgr.ClientSecret != cli.GetSecret() {
		return nil, errors.ErrInvalidClient
	}

	if gt == definition.ClientCredentials && cli.IsPublic() == true {
		return nil, errors.ErrInvalidClient
	}

	if gt == definition.AuthorizationCode {
		ci, err := m.getAndDelAuthorizationCode(ctx, tgr)
		if err != nil {
			return nil, err
		}
		if err := m.validateCodeChallenge(ci, tgr.CodeVerifier); err != nil {
			return nil, err
		}
		tgr.UserID = ci.GetUserID()
		tgr.Scope = ci.GetScope()
	}

	gcfg := m.grantConfig(gt)
	createAt := time.Now()
	td := &definition.GenerateBasic{
		Client:    cli,
		UserID:    tgr.UserID,
		Scope:     tgr.Scope,
		CreatedAt: createAt,
	}
	body, err := m.accessGenerate.Token(ctx, td, gcfg.IsGenerateRefresh)
	if err != nil {
		return nil, err
	}

	ti := m.token.New()
	ti.SetClientID(tgr.ClientID)
	ti.SetUserID(tgr.UserID)
	ti.SetScope(tgr.Scope)
	ti.SetExpiredAt(createAt.Add(gcfg.AccessTokenExp))
	td.ExpiredAt = ti.GetExpiredAt()
	ti.SetAccess(body.Access)
	ti.SetAccessIdentifier(body.AccessIdentifier)
	if gcfg.IsGenerateRefresh && body.Refresh != "" {
		ti.SetRefresh(body.Refresh)
		ti.SetRefreshExpiredAt(createAt.Add(gcfg.RefreshTokenExp))
	}

	err = m.tokenStore.Create(ctx, ti)
	if err != nil {
		return nil, err
	}

	return ti, nil
}

// RefreshAccessToken refreshing an access token
func (m *Manager) RefreshAccessToken(ctx context.Context, tgr *definition.TokenGenerateRequest) (definition.TokenInfo, error) {
	ti, err := m.LoadRefreshToken(ctx, tgr.Refresh)
	if err != nil {
		return nil, err
	}

	cli, err := m.GetClient(ctx, ti.GetClientID())
	if err != nil {
		return nil, err
	}

	oldAccess, oldRefresh := ti.GetAccess(), ti.GetRefresh()

	td := &definition.GenerateBasic{
		Client:    cli,
		UserID:    ti.GetUserID(),
		Scope:     ti.GetScope(),
		CreatedAt: time.Now(),
		Request:   tgr.Request,
	}

	rcfg := DefaultRefreshTokenCfg
	if v := m.rcfg; v != nil {
		rcfg = v
	}

	if rcfg.AccessTokenExp > 0 {
		ti.SetExpiredAt(time.Now().Add(rcfg.AccessTokenExp))
		td.ExpiredAt = ti.GetExpiredAt()
	}
	if rcfg.IsResetRefreshTime && rcfg.RefreshTokenExp > 0 {
		ti.SetRefreshExpiredAt(time.Now().Add(rcfg.RefreshTokenExp))
	}

	if scope := tgr.Scope; scope != "" {
		td.Scope = scope
		ti.SetScope(scope)
	}

	body, err := m.accessGenerate.Token(ctx, td, rcfg.IsGenerateRefresh)
	if err != nil {
		return nil, err
	}

	ti.SetAccess(body.Access)
	ti.SetAccessIdentifier(body.AccessIdentifier)
	if body.Refresh != "" {
		ti.SetRefresh(body.Refresh)
	}

	if err := m.tokenStore.Create(ctx, ti); err != nil {
		return nil, err
	}

	if rcfg.IsRemoveAccess {
		// remove the old access token
		if err := m.tokenStore.RemoveByAccess(ctx, oldAccess); err != nil {
			return nil, err
		}
	}

	if rcfg.IsRemoveRefreshing && body.Refresh != "" {
		// remove the old refresh token
		if err := m.tokenStore.RemoveByRefresh(ctx, oldRefresh); err != nil {
			return nil, err
		}
	}

	return ti, nil
}

// RemoveAccessToken use the access token to delete the token information
func (m *Manager) RemoveAccessToken(ctx context.Context, access string) error {
	if access == "" {
		return errors.ErrInvalidAccessToken
	}
	return m.tokenStore.RemoveByAccess(ctx, access)
}

// RemoveRefreshToken use the refresh token to delete the token information
func (m *Manager) RemoveRefreshToken(ctx context.Context, refresh string) error {
	if refresh == "" {
		return errors.ErrInvalidAccessToken
	}
	return m.tokenStore.RemoveByRefresh(ctx, refresh)
}

// LoadAccessToken according to the access token for corresponding token information
func (m *Manager) LoadAccessToken(ctx context.Context, access string) (definition.TokenInfo, error) {
	if access == "" {
		return nil, errors.ErrInvalidAccessToken
	}

	ct := time.Now()
	ti, err := m.tokenStore.GetByAccess(ctx, access)
	if err != nil {
		return nil, err
	} else if ti == nil {
		return nil, errors.ErrInvalidAccessToken
	} else if ti.GetExpiredAt().Before(ct) {
		return nil, errors.ErrExpiredAccessToken
	}
	return ti, nil
}

// LoadRefreshToken according to the refresh token for corresponding token information
func (m *Manager) LoadRefreshToken(ctx context.Context, refresh string) (definition.TokenInfo, error) {
	if refresh == "" {
		return nil, errors.ErrInvalidRefreshToken
	}

	ti, err := m.tokenStore.GetByRefresh(ctx, refresh)
	if err != nil {
		return nil, err
	} else if ti == nil {
		return nil, errors.ErrInvalidRefreshToken
	} else if ti.GetRefreshExpiredAt().Before(time.Now()) {
		return nil, errors.ErrExpiredRefreshToken
	}
	return ti, nil
}
