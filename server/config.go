package server

import (
	"net/http"
	"time"

	"github.com/RandalTeng/oauth2/definition"
)

// Config configuration parameters
type Config struct {
	TokenType                   string                    // token type
	AllowGetAccessRequest       bool                      // to allow GET requests for the token
	AllowedResponseTypes        []definition.ResponseType // allow the authorization type
	AllowedGrantTypes           []definition.GrantType    // allow the grant type
	AllowedCodeChallengeMethods []definition.CodeChallengeMethod
	ForcePKCE                   bool
}

// NewConfig create to configuration instance
func NewConfig() *Config {
	return &Config{
		TokenType:            "Bearer",
		AllowedResponseTypes: []definition.ResponseType{definition.Code, definition.Token},
		AllowedGrantTypes: []definition.GrantType{
			definition.AuthorizationCode,
			definition.PasswordCredentials,
			definition.ClientCredentials,
			definition.Refreshing,
		},
		AllowedCodeChallengeMethods: []definition.CodeChallengeMethod{
			definition.CodeChallengePlain,
			definition.CodeChallengeS256,
		},
	}
}

// AuthorizeRequest authorization request
type AuthorizeRequest struct {
	ResponseType        definition.ResponseType
	ClientID            string
	Scope               string
	RedirectURI         string
	State               string
	UserID              string
	CodeChallenge       string
	CodeChallengeMethod definition.CodeChallengeMethod
	AccessTokenExp      time.Duration
	Request             *http.Request
}
