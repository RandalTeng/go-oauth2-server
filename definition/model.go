package definition

import "time"

type (
	// ClientInfo the client information model interface
	ClientInfo interface {
		GetID() string
		GetSecret() string
		GetDomain() string
		IsPublic() bool
		GetUserID() string
	}

	// ClientPasswordVerifier the password handler interface
	ClientPasswordVerifier interface {
		VerifyPassword(string) bool
	}

	// CodeInfo the authorization code information model interface
	CodeInfo interface {
		New() CodeInfo

		GetClientID() string
		SetClientID(string)
		GetUserID() string
		SetUserID(string)
		GetRedirectURI() string
		SetRedirectURI(string)
		GetScope() string
		SetScope(string)

		GetCode() string
		SetCode(string)
		GetCodeCreateAt() time.Time
		SetCodeCreateAt(time.Time)
		GetCodeExpiresIn() time.Duration
		SetCodeExpiresIn(time.Duration)
		GetCodeChallenge() string
		SetCodeChallenge(string)
		GetCodeChallengeMethod() CodeChallengeMethod
		SetCodeChallengeMethod(CodeChallengeMethod)
	}

	// TokenInfo the access token information model interface
	TokenInfo interface {
		New() TokenInfo

		GetClientID() string
		SetClientID(string)
		GetUserID() string
		SetUserID(string)
		GetScope() string
		SetScope(string)

		GetAccess() string
		SetAccess(string)
		GetAccessCreateAt() time.Time
		SetAccessCreateAt(time.Time)
		GetAccessExpiresIn() time.Duration
		SetAccessExpiresIn(time.Duration)
	}

	// RefreshInfo the refresh token information model interface
	RefreshInfo interface {
		New() RefreshInfo

		GetRefresh() string
		SetRefresh(string)
		GetRefreshCreateAt() time.Time
		SetRefreshCreateAt(time.Time)
		GetRefreshExpiresIn() time.Duration
		SetRefreshExpiresIn(time.Duration)
	}
)
