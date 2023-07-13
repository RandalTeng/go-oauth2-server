package generator

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/RandalTeng/go-oauth2-server/definition"
	"github.com/RandalTeng/go-oauth2-server/models"
	"github.com/golang-jwt/jwt/v4"
)

func TestJWTAccessGenerate_Token(t *testing.T) {
	// Just for hs256 string sign verify, sign string validate should have completed
	// on JWT package.
	secret := bytes.NewBufferString("1234567890abcdef").Bytes()
	generator := NewJWTAccessGenerate("tester", secret, jwt.SigningMethodHS256)
	basic := &definition.GenerateBasic{
		Client:    &models.Client{ID: "abc"},
		UserID:    "tester",
		Scope:     "scope1 scope2",
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(2 * time.Hour),
	}
	body, err := generator.Token(context.Background(), basic, false)
	if err != nil {
		t.Errorf("generate random access token error: %+v", err)
	} else if len(body.Access) == 0 {
		t.Errorf("generate random access token error: access len is 0")
	} else if len(body.AccessIdentifier) == 0 {
		t.Errorf("generate random access token error: access identify len is 0, JWT token should have a identify string")
	}
	body, err = generator.Token(context.Background(), basic, true)
	if err != nil {
		t.Errorf("generate random access token error: %+v", err)
	} else if len(body.Refresh) == 0 {
		t.Errorf("generate random access token error: refressh token len is 0")
	}
}
