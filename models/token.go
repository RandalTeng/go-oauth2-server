package models

import (
	"time"

	"github.com/RandalTeng/go-oauth2-server/definition"
)

// NewToken create to token model instance
func NewToken() *Token {
	return &Token{}
}

// Token token model
type Token struct {
	ClientID         string
	TokenType        string
	UserID           string
	Scope            string
	Access           string
	AccessIdentifier string
	ExpiredAt        time.Time
	Refresh          string
	RefreshExpiredAt time.Time
}

// New create to token model instance
func (t *Token) New() definition.TokenInfo {
	return NewToken()
}

// GetClientID the client id
func (t *Token) GetClientID() string {
	return t.ClientID
}

// SetClientID the client id
func (t *Token) SetClientID(clientID string) {
	t.ClientID = clientID
}

// GetUserID the user id
func (t *Token) GetUserID() string {
	return t.UserID
}

// SetUserID the user id
func (t *Token) SetUserID(userID string) {
	t.UserID = userID
}

// GetScope get scope of authorization
func (t *Token) GetScope() string {
	return t.Scope
}

// SetScope get scope of authorization
func (t *Token) SetScope(scope string) {
	t.Scope = scope
}

// GetAccess access Token
func (t *Token) GetAccess() string {
	return t.Access
}

// SetAccess access Token
func (t *Token) SetAccess(access string) {
	t.Access = access
}

// GetAccessIdentifier access id
func (t *Token) GetAccessIdentifier() string {
	return t.AccessIdentifier
}

// SetAccessIdentifier access id
func (t *Token) SetAccessIdentifier(access string) {
	t.AccessIdentifier = access
}

// GetExpiredAt the lifetime in seconds of the access token
func (t *Token) GetExpiredAt() time.Time {
	return t.ExpiredAt
}

// SetExpiredAt the lifetime in seconds of the access token
func (t *Token) SetExpiredAt(exp time.Time) {
	t.ExpiredAt = exp
}

// GetRefresh refresh Token
func (t *Token) GetRefresh() string {
	return t.Refresh
}

// SetRefresh refresh Token
func (t *Token) SetRefresh(refresh string) {
	t.Refresh = refresh
}

// GetRefreshExpiredAt the lifetime in seconds of the refresh token
func (t *Token) GetRefreshExpiredAt() time.Time {
	return t.RefreshExpiredAt
}

// SetRefreshExpiredAt the lifetime in seconds of the refresh token
func (t *Token) SetRefreshExpiredAt(exp time.Time) {
	t.RefreshExpiredAt = exp
}
