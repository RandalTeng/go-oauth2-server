package generator

import (
	"context"
	"testing"

	"github.com/RandalTeng/go-oauth2-server/definition"
	"github.com/RandalTeng/go-oauth2-server/models"
)

func TestAuthorizeGenerate_Token(t *testing.T) {
	generator := NewAuthorizeGenerate()
	basic := &definition.GenerateBasic{
		Client: &models.Client{ID: "abc"},
		UserID: "tester",
	}
	code, err := generator.Token(context.Background(), basic)
	if err != nil {
		t.Errorf("generate authorize code error: %v", err)
	} else if len(code) == 0 {
		t.Errorf("generate authorize code error: code len is 0")
	}
}
