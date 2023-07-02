package generator

import (
	"context"
	"strings"
	"time"

	"github.com/RandalTeng/go-oauth2-server/definition"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// JWTAccessClaims jwt claims
type JWTAccessClaims struct {
	jwt.RegisteredClaims
	Scope string `json:"scp,omitempty"`
}

// NewJWTAccessGenerate create to generate the jwt access token instance
func NewJWTAccessGenerate(kid string, key []byte, method jwt.SigningMethod) *JWTAccessGenerate {
	return &JWTAccessGenerate{
		SignedKeyID:  kid,
		SignedKey:    key,
		SignedMethod: method,
	}
}

// JWTAccessGenerate generate the jwt access token
type JWTAccessGenerate struct {
	SignedKeyID  string
	SignedKey    []byte
	SignedMethod jwt.SigningMethod
}

// Token based on the UUID generated token
func (a *JWTAccessGenerate) Token(ctx context.Context, data *definition.GenerateBasic, isGenRefresh bool) (body definition.GenerateTokenBody, err error) {
	body.AccessIdentifier = uuidString()
	claims := &JWTAccessClaims{
		Scope: data.TokenInfo.GetScope(),
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        body.AccessIdentifier,
			Audience:  jwt.ClaimStrings{data.Client.GetID()},
			Subject:   data.UserID,
			ExpiresAt: jwt.NewNumericDate(data.TokenInfo.GetExpiredAt()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(a.SignedMethod, claims)
	if a.SignedKeyID != "" {
		token.Header["kid"] = a.SignedKeyID
	}
	var key interface{}
	if a.isEs() {
		key, err = jwt.ParseECPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return
		}
	} else if a.isRsOrPS() {
		key, err = jwt.ParseRSAPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return
		}
	} else if a.isHs() {
		key = a.SignedKey
	} else if a.isEd() {
		key, err = jwt.ParseEdPrivateKeyFromPEM(a.SignedKey)
		if err != nil {
			return
		}
	} else {
		return
	}

	body.Access, err = token.SignedString(key)
	if err != nil {
		return
	}
	if isGenRefresh {
		body.Refresh = uuidString()
	}

	return
}

func (a *JWTAccessGenerate) isEs() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "ES")
}

func (a *JWTAccessGenerate) isRsOrPS() bool {
	isRs := strings.HasPrefix(a.SignedMethod.Alg(), "RS")
	isPs := strings.HasPrefix(a.SignedMethod.Alg(), "PS")
	return isRs || isPs
}

func (a *JWTAccessGenerate) isHs() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "HS")
}

func (a *JWTAccessGenerate) isEd() bool {
	return strings.HasPrefix(a.SignedMethod.Alg(), "Ed")
}

// generate uuid string without "-".
func uuidString() string {
	return strings.Replace(uuid.NewString(), "-", "", -1)
}
