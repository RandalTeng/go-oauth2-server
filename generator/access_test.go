package generator

import (
	"context"
	"testing"
	"time"

	"github.com/RandalTeng/go-oauth2-server/definition"
	"github.com/RandalTeng/go-oauth2-server/models"
)

func TestAccessGenerate_Token(t *testing.T) {
	generator := NewAccessGenerate()
	basic := &definition.GenerateBasic{
		Client:    &models.Client{ID: "abc"},
		UserID:    "tester",
		CreatedAt: time.Now(),
	}
	body, err := generator.Token(context.Background(), basic, false)
	if err != nil {
		t.Errorf("generate random access token error: %+v", err)
	} else if len(body.Access) == 0 {
		t.Errorf("generate random access token error: access len is 0")
	}
	body, err = generator.Token(context.Background(), basic, true)
	if err != nil {
		t.Errorf("generate random access token error: %+v", err)
	} else if len(body.Refresh) == 0 {
		t.Errorf("generate random access token error: refressh token len is 0")
	}
}
