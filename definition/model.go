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
		CodeInfoAdapter

		GetClientID() string
		SetClientID(string)
		GetUserID() string
		SetUserID(string)
		GetRedirectURI() string
		SetRedirectURI(string)
		GetScope() string
		SetScope(string)

		GetCode() string
		// SetCode code should be a random string that 32 length size, seem like a string id.
		SetCode(string)
		GetChallenge() string
		SetChallenge(string)
		GetChallengeMethod() CodeChallengeMethod
		SetChallengeMethod(CodeChallengeMethod)
		GetExpiredAt() time.Time
		SetExpiredAt(time.Time)
		// CreatedAt is optional, will not force package user set created at.
		// Revoked is optional, will not force package user delete some auth code hard.
	}

	// TokenInfo the access token information model interface
	TokenInfo interface {
		TokenInfoAdapter

		GetClientID() string
		SetClientID(string)
		GetUserID() string
		SetUserID(string)
		GetScope() string
		SetScope(string)

		GetAccess() string
		// SetAccess access should be a random string that 32 length size, seem like a string id.
		SetAccess(string)
		GetAccessIdentifier() string
		SetAccessIdentifier(string)
		GetExpiredAt() time.Time
		SetExpiredAt(time.Time)
		GetRefresh() string
		// SetRefresh refresh should be a random string that 32 length size, seem like a string id.
		SetRefresh(string)
		GetRefreshExpiredAt() time.Time
		SetRefreshExpiredAt(time.Time)
		// CreatedAt is optional, will not force package user set created at.
		// Revoked is optional, will not force package user delete some auth code hard.
	}

	CodeInfoAdapter interface {
		New() CodeInfo
	}

	TokenInfoAdapter interface {
		New() TokenInfo
	}
)
