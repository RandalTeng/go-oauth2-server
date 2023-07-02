package models

import (
	"time"

	"github.com/RandalTeng/go-oauth2-server/definition"
)

// NewCode create to code model instance
func NewCode() *Code {
	return &Code{}
}

// Code model
type Code struct {
	ClientID        string
	UserID          string
	RedirectURI     string
	Scope           string
	Code            string
	Challenge       string
	ChallengeMethod string
	ExpiredAt       time.Time
}

// New create to token model instance
func (t *Code) New() definition.CodeInfo {
	return NewCode()
}

// GetClientID the client id
func (t *Code) GetClientID() string {
	return t.ClientID
}

// SetClientID the client id
func (t *Code) SetClientID(clientID string) {
	t.ClientID = clientID
}

// GetUserID the user id
func (t *Code) GetUserID() string {
	return t.UserID
}

// SetUserID the user id
func (t *Code) SetUserID(userID string) {
	t.UserID = userID
}

// GetRedirectURI redirect URI
func (t *Code) GetRedirectURI() string {
	return t.RedirectURI
}

// SetRedirectURI redirect URI
func (t *Code) SetRedirectURI(redirectURI string) {
	t.RedirectURI = redirectURI
}

// GetScope get scope of authorization
func (t *Code) GetScope() string {
	return t.Scope
}

// SetScope get scope of authorization
func (t *Code) SetScope(scope string) {
	t.Scope = scope
}

// GetCode authorization code
func (t *Code) GetCode() string {
	return t.Code
}

// SetCode authorization code
func (t *Code) SetCode(code string) {
	t.Code = code
}

// GetChallenge challenge code
func (t *Code) GetChallenge() string {
	return t.Challenge
}

// SetChallenge challenge code
func (t *Code) SetChallenge(code string) {
	t.Challenge = code
}

// GetChallengeMethod challenge method
func (t *Code) GetChallengeMethod() definition.CodeChallengeMethod {
	return definition.CodeChallengeMethod(t.ChallengeMethod)
}

// SetChallengeMethod challenge method
func (t *Code) SetChallengeMethod(method definition.CodeChallengeMethod) {
	t.ChallengeMethod = string(method)
}

// GetExpiredAt the lifetime in seconds of the authorization code
func (t *Code) GetExpiredAt() time.Time {
	return t.ExpiredAt
}

// SetExpiredAt the lifetime in seconds of the authorization code
func (t *Code) SetExpiredAt(exp time.Time) {
	t.ExpiredAt = exp
}
