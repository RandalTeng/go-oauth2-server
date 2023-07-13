package manager

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/RandalTeng/go-oauth2-server/definition"
	"github.com/RandalTeng/go-oauth2-server/models"
	"github.com/RandalTeng/go-oauth2-server/store"
)

func TestManager_GenerateAccessToken(t *testing.T) {
	cliStrg := store.NewClientStore()
	testClient := &models.Client{ID: "1001", Secret: "123456", Domain: "https://example.com/", UserID: "1001"}
	err := cliStrg.Set(testClient.GetID(), testClient)
	if err != nil {
		t.Fatalf("Create test client error: %v", err)
	}

	tokenStrg := store.NewMemoryTokenStore(3)

	m := NewDefaultManager()
	m.MapClientStorage(cliStrg)
	m.MapTokenStorage(tokenStrg)
	m.SetTokenAdapter(models.NewToken())
	type args struct {
		gt  definition.GrantType
		tgr *definition.TokenGenerateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    definition.TokenInfo
		wantErr bool
	}{
		{
			name: "client credentials success",
			args: args{gt: definition.ClientCredentials, tgr: &definition.TokenGenerateRequest{
				ClientID:     "1001",
				ClientSecret: "123456",
				UserID:       "10001",
				Scope:        "scope1 scope2",
			}},
			want:    nil,
			wantErr: false,
		},
		{
			name: "client credentials failure",
			args: args{gt: definition.ClientCredentials, tgr: &definition.TokenGenerateRequest{
				ClientID:     "1001",
				ClientSecret: "",
				UserID:       "10001",
				Scope:        "scope1 scope2",
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GenerateAccessToken(context.Background(), tt.args.gt, tt.args.tgr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.GetClientID() != testClient.GetID() {
				t.Errorf("GenerateAccessToken() got clientId: %s, want clientId: %s", got.GetClientID(), testClient.GetID())
			}
		})
	}
}

func TestManager_GenerateAuthToken(t *testing.T) {
	cliStrg := store.NewClientStore()
	testClient := &models.Client{ID: "1001", Secret: "123456", Domain: "https://example.com/", UserID: "1001"}
	err := cliStrg.Set(testClient.GetID(), testClient)
	if err != nil {
		t.Fatalf("Create test client error: %v", err)
	}

	tokenStrg := store.NewMemoryTokenStore(3)

	m := NewDefaultManager()
	m.MapClientStorage(cliStrg)
	m.MapTokenStorage(tokenStrg)
	m.SetTokenAdapter(models.NewToken())
	type args struct {
		rt  definition.ResponseType
		tgr *definition.TokenGenerateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "code success",
			args: args{rt: definition.Code, tgr: &definition.TokenGenerateRequest{}},
		},
		{
			name: "code failure",
			args: args{rt: definition.Code, tgr: &definition.TokenGenerateRequest{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GenerateAuthToken(context.Background(), tt.args.rt, tt.args.tgr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateAuthToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_GetClient(t *testing.T) {
	cliStrg := store.NewClientStore()
	testClient := &models.Client{ID: "1001", Secret: "123456", Domain: "https://example.com/", UserID: "1001"}
	err := cliStrg.Set(testClient.GetID(), testClient)
	if err != nil {
		t.Fatalf("Create test client error: %v", err)
	}
	m := NewDefaultManager()
	m.MapClientStorage(cliStrg)

	type args struct {
		cid string
	}
	tests := []struct {
		name    string
		args    args
		want    definition.ClientInfo
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{cid: testClient.GetID()},
			want:    testClient,
			wantErr: false,
		},
		{
			name:    "failure",
			args:    args{cid: "1002"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetClient(context.Background(), tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadAccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_LoadAccessToken(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ctx    context.Context
		access string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    definition.TokenInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			got, err := m.LoadAccessToken(tt.args.ctx, tt.args.access)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadAccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_LoadRefreshToken(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ctx     context.Context
		refresh string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    definition.TokenInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			got, err := m.LoadRefreshToken(tt.args.ctx, tt.args.refresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadRefreshToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_MapAccessGenerate(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		gen definition.AccessGenerate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.MapAccessGenerate(tt.args.gen)
		})
	}
}

func TestManager_MapAuthorizeGenerate(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		gen definition.AuthorizeGenerate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.MapAuthorizeGenerate(tt.args.gen)
		})
	}
}

func TestManager_MapClientStorage(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		stor definition.ClientStore
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.MapClientStorage(tt.args.stor)
		})
	}
}

func TestManager_MapTokenStorage(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		stor definition.TokenStore
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.MapTokenStorage(tt.args.stor)
		})
	}
}

func TestManager_MustClientStorage(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		stor definition.ClientStore
		err  error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.MustClientStorage(tt.args.stor, tt.args.err)
		})
	}
}

func TestManager_MustTokenStorage(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		stor definition.TokenStore
		err  error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.MustTokenStorage(tt.args.stor, tt.args.err)
		})
	}
}

func TestManager_RefreshAccessToken(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ctx context.Context
		tgr *definition.TokenGenerateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    definition.TokenInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			got, err := m.RefreshAccessToken(tt.args.ctx, tt.args.tgr)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefreshAccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_RemoveAccessToken(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ctx    context.Context
		access string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			if err := m.RemoveAccessToken(tt.args.ctx, tt.args.access); (err != nil) != tt.wantErr {
				t.Errorf("RemoveAccessToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_RemoveRefreshToken(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ctx     context.Context
		refresh string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			if err := m.RemoveRefreshToken(tt.args.ctx, tt.args.refresh); (err != nil) != tt.wantErr {
				t.Errorf("RemoveRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_SetAuthorizeCodeExp(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		exp time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.SetAuthorizeCodeExp(tt.args.exp)
		})
	}
}

func TestManager_SetTokenCfgWithSpecificType(t *testing.T) {
	m := NewDefaultManager()
	type args struct {
		gt  definition.GrantType
		cfg *Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "authorize code",
			args: args{gt: definition.AuthorizationCode, cfg: DefaultAuthorizeCodeTokenCfg},
		},
		{
			name: "password",
			args: args{gt: definition.PasswordCredentials, cfg: DefaultPasswordTokenCfg},
		},
		{
			name: "client credentials",
			args: args{gt: definition.ClientCredentials, cfg: DefaultClientTokenCfg},
		},
		{
			name: "implicit",
			args: args{gt: definition.Implicit, cfg: DefaultImplicitTokenCfg},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.SetTokenCfgWithType(tt.args.gt, tt.args.cfg)
		})
	}
}

func TestManager_SetCodeAdapter(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		adp definition.CodeInfoAdapter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.SetCodeAdapter(tt.args.adp)
		})
	}
}

func TestManager_SetLogger(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		logger definition.Logger
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.SetLogger(tt.args.logger)
		})
	}
}

func TestManager_SetRefreshTokenCfg(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		cfg *RefreshingConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.SetRefreshTokenCfg(tt.args.cfg)
		})
	}
}

func TestManager_SetTokenAdapter(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		adp definition.TokenInfoAdapter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.SetTokenAdapter(tt.args.adp)
		})
	}
}

func TestManager_SetValidateURIHandler(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		handler ValidateURIHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.SetValidateURIHandler(tt.args.handler)
		})
	}
}

func TestManager_delAuthorizationCode(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			if err := m.delAuthorizationCode(tt.args.ctx, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("delAuthorizationCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_getAndDelAuthorizationCode(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ctx context.Context
		tgr *definition.TokenGenerateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    definition.CodeInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			got, err := m.getAndDelAuthorizationCode(tt.args.ctx, tt.args.tgr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAndDelAuthorizationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAndDelAuthorizationCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_getAuthorizationCode(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    definition.CodeInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			got, err := m.getAuthorizationCode(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAuthorizationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAuthorizationCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_grantConfig(t *testing.T) {
	m := NewDefaultManager()
	type args struct {
		gt definition.GrantType
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "authorize code",
			args: args{gt: definition.AuthorizationCode},
			want: DefaultAuthorizeCodeTokenCfg,
		},
		{
			name: "password",
			args: args{gt: definition.PasswordCredentials},
			want: DefaultPasswordTokenCfg,
		},
		{
			name: "client credentials",
			args: args{gt: definition.ClientCredentials},
			want: DefaultClientTokenCfg,
		},
		{
			name: "implicit",
			args: args{gt: definition.Implicit},
			want: DefaultImplicitTokenCfg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m.grantConfig(tt.args.gt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grantConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_log(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		level  definition.LogLevel
		format string
		args   []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			m.log(tt.args.level, tt.args.format, tt.args.args...)
		})
	}
}

func TestManager_validateCodeChallenge(t *testing.T) {
	type fields struct {
		codeExp           time.Duration
		gtcfg             map[definition.GrantType]*Config
		rcfg              *RefreshingConfig
		validateURI       ValidateURIHandler
		authorizeGenerate definition.AuthorizeGenerate
		accessGenerate    definition.AccessGenerate
		tokenStore        definition.TokenStore
		clientStore       definition.ClientStore
		logger            definition.Logger
		code              definition.CodeInfoAdapter
		token             definition.TokenInfoAdapter
	}
	type args struct {
		ci  definition.CodeInfo
		ver string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				codeExp:           tt.fields.codeExp,
				gtcfg:             tt.fields.gtcfg,
				rcfg:              tt.fields.rcfg,
				validateURI:       tt.fields.validateURI,
				authorizeGenerate: tt.fields.authorizeGenerate,
				accessGenerate:    tt.fields.accessGenerate,
				tokenStore:        tt.fields.tokenStore,
				clientStore:       tt.fields.clientStore,
				logger:            tt.fields.logger,
				code:              tt.fields.code,
				token:             tt.fields.token,
			}
			if err := m.validateCodeChallenge(tt.args.ci, tt.args.ver); (err != nil) != tt.wantErr {
				t.Errorf("validateCodeChallenge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDefaultManager(t *testing.T) {
	tests := []struct {
		name string
		want *Manager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewManager(t *testing.T) {
	tests := []struct {
		name string
		want *Manager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
