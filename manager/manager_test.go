package manager

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/RandalTeng/go-oauth2-server/definition"
	"github.com/RandalTeng/go-oauth2-server/models"
	"github.com/RandalTeng/go-oauth2-server/store"
)

func TestManager_GenerateAccessToken(t *testing.T) {
	cliStrg := store.NewClientStore()
	testClient := &models.Client{ID: "1001", Secret: "123456", Domain: "https://example.com/", UserID: "1001"}
	err := cliStrg.Set(testClient.GetID(), testClient)
	if err != nil {
		t.Fatalf("Create test client error: %v", err)
	}

	tokenStrg := store.NewMemoryTokenStore(3)

	m := NewDefaultManager()
	m.MapClientStorage(cliStrg)
	m.MapTokenStorage(tokenStrg)
	m.SetTokenAdapter(models.NewToken())
	type args struct {
		gt  definition.GrantType
		tgr *definition.TokenGenerateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    definition.TokenInfo
		wantErr bool
	}{
		{
			name: "client credentials success",
			args: args{gt: definition.ClientCredentials, tgr: &definition.TokenGenerateRequest{
				ClientID:     "1001",
				ClientSecret: "123456",
				UserID:       "10001",
				Scope:        "scope1 scope2",
			}},
			want:    nil,
			wantErr: false,
		},
		{
			name: "client credentials failure",
			args: args{gt: definition.ClientCredentials, tgr: &definition.TokenGenerateRequest{
				ClientID:     "1001",
				ClientSecret: "",
				UserID:       "10001",
				Scope:        "scope1 scope2",
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GenerateAccessToken(context.Background(), tt.args.gt, tt.args.tgr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.GetClientID() != testClient.GetID() {
				t.Errorf("GenerateAccessToken() got clientId: %s, want clientId: %s", got.GetClientID(), testClient.GetID())
			}
		})
	}
}

func TestManager_GenerateAuthToken(t *testing.T) {
	cliStrg := store.NewClientStore()
	testClient := &models.Client{ID: "1001", Secret: "123456", Domain: "https://example.com/", UserID: "1001"}
	err := cliStrg.Set(testClient.GetID(), testClient)
	if err != nil {
		t.Fatalf("Create test client error: %v", err)
	}

	tokenStrg := store.NewMemoryTokenStore(3)

	m := NewDefaultManager()
	m.MapClientStorage(cliStrg)
	m.MapTokenStorage(tokenStrg)
	m.SetCodeAdapter(models.NewCode())
	m.SetTokenAdapter(models.NewToken())
	type args struct {
		rt  definition.ResponseType
		tgr *definition.TokenGenerateRequest
	}
	codeSuccessReq, _ := http.NewRequest("GET", "", strings.NewReader(""))
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "code success",
			args: args{rt: definition.Code, tgr: &definition.TokenGenerateRequest{
				ClientID:    "1001",
				RedirectURI: "https://example.com/oauth/callback",
				UserID:      "10001",
				Scope:       "scope1 scope2",
				Request:     codeSuccessReq,
			}},
			wantErr: false,
		},
		{
			name: "code failure",
			args: args{rt: definition.Code, tgr: &definition.TokenGenerateRequest{
				ClientID:    "1002",
				RedirectURI: "https://example-test.com/oauth/callback",
				UserID:      "10001",
				Scope:       "scope1 scope2",
				Request:     codeSuccessReq,
			}},
			wantErr: true,
		},
		{
			name: "token success",
			args: args{rt: definition.Token, tgr: &definition.TokenGenerateRequest{
				ClientID:    "1001",
				RedirectURI: "https://example.com/oauth/callback",
				UserID:      "10001",
				Scope:       "scope1 scope2",
				Request:     codeSuccessReq,
			}},
			wantErr: false,
		},
		{
			name: "token failure",
			args: args{rt: definition.Token, tgr: &definition.TokenGenerateRequest{
				ClientID:    "1002",
				RedirectURI: "https://example-test.com/oauth/callback",
				UserID:      "10001",
				Scope:       "scope1 scope2",
				Request:     codeSuccessReq,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = m.GenerateAuthToken(context.Background(), tt.args.rt, tt.args.tgr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestManager_RefreshAccessToken(t *testing.T) {
	cliStrg := store.NewClientStore()
	testClient := &models.Client{ID: "1001", Secret: "123456", Domain: "https://example.com/", UserID: "1001"}
	err := cliStrg.Set(testClient.GetID(), testClient)
	if err != nil {
		t.Fatalf("Create test client error: %v", err)
	}

	tokenStrg := store.NewMemoryTokenStore(3)
	_ = tokenStrg.Create(context.Background(), &models.Token{
		ClientID:         "1001",
		UserID:           "10001",
		Scope:            "scope1 scope1",
		Access:           "abc123",
		AccessIdentifier: "abc123",
		Refresh:          "abcdef",
		RefreshExpiredAt: time.Date(2033, 7, 17, 0, 0, 0, 0, time.Local),
	})

	m := NewDefaultManager()
	m.MapClientStorage(cliStrg)
	m.MapTokenStorage(tokenStrg)
	m.SetCodeAdapter(models.NewCode())
	m.SetTokenAdapter(models.NewToken())

	type args struct {
		tgr *definition.TokenGenerateRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{tgr: &definition.TokenGenerateRequest{
				Refresh: "abcdef",
			}},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{tgr: &definition.TokenGenerateRequest{
				Refresh: "abcdefg",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = m.RefreshAccessToken(context.Background(), tt.args.tgr)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
