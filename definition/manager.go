package definition

import (
	"context"
	"net/http"
)

// TokenGenerateRequest provide to generate the token request parameters
type TokenGenerateRequest struct {
	ClientID            string
	ClientSecret        string
	UserID              string
	RedirectURI         string
	Scope               string
	Code                string
	CodeChallenge       string
	CodeChallengeMethod CodeChallengeMethod
	Refresh             string
	CodeVerifier        string
	Request             *http.Request
}

// Manager authorization management interface
type Manager interface {
	// GetClient get the client information
	GetClient(ctx context.Context, clientID string) (cli ClientInfo, err error)

	// SetCodeAdapter auth code model adapter
	SetCodeAdapter(adp CodeInfoAdapter)
	// SetTokenAdapter access token model adapter
	SetTokenAdapter(adp TokenInfoAdapter)

	// GenerateAuthToken generate the authorization code
	GenerateAuthToken(ctx context.Context, rt ResponseType, tgr *TokenGenerateRequest) (rb any, err error)

	// GenerateAccessToken generate the access token(authorize token)
	GenerateAccessToken(ctx context.Context, gt GrantType, tgr *TokenGenerateRequest) (accessToken TokenInfo, err error)

	// RefreshAccessToken refreshing an access token
	RefreshAccessToken(ctx context.Context, tgr *TokenGenerateRequest) (accessToken TokenInfo, err error)

	// RemoveAccessToken use the access token to delete the token information
	RemoveAccessToken(ctx context.Context, access string) (err error)

	// RemoveRefreshToken use the refresh token to delete the token information
	RemoveRefreshToken(ctx context.Context, refresh string) (err error)

	// LoadAccessToken according to the access token for corresponding token information
	LoadAccessToken(ctx context.Context, access string) (ti TokenInfo, err error)

	// LoadRefreshToken according to the refresh token for corresponding token information
	LoadRefreshToken(ctx context.Context, refresh string) (ti TokenInfo, err error)
}
