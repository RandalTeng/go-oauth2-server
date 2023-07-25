package server

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"
)

func TestClientBasicHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}
	basic := base64.StdEncoding.EncodeToString(bytes.NewBufferString("test1:pwd1").Bytes())
	r1, _ := http.NewRequest(http.MethodGet, "", bytes.NewBufferString(""))
	r1.Header.Set("Authorization", fmt.Sprintf("Basic %s", basic))
	tests := []struct {
		name     string
		args     args
		username string
		password string
		wantErr  bool
	}{
		{
			name:     "success",
			args:     args{r: r1},
			username: "test1",
			password: "pwd1",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			username, password, err := ClientBasicHandler(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientBasicHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if username != tt.username {
				t.Errorf("ClientBasicHandler() got = %v, username %v", username, tt.username)
			}
			if password != tt.password {
				t.Errorf("ClientBasicHandler() got1 = %v, username %v", password, tt.password)
			}
		})
	}
}

func TestClientBodyHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}
	r1, _ := http.NewRequest(http.MethodPost, "", bytes.NewBufferString(`{"client_id":"client1","client_secret":"secret1"}`))
	r1.Header.Set("Content-Type", "application/json")
	r2, _ := http.NewRequest(http.MethodPost, "", bytes.NewBufferString(`client_id=client2&client_secret=secret2`))
	r2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_ = r2.ParseForm()
	tests := []struct {
		name     string
		args     args
		clientId string
		secret   string
		wantErr  bool
	}{
		{
			name:     "json",
			args:     args{r: r1},
			clientId: "client1",
			secret:   "secret1",
			wantErr:  false,
		},
		{
			name:     "form",
			args:     args{r: r2},
			clientId: "client2",
			secret:   "secret2",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientId, secret, err := ClientBodyHandler(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientBodyHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if clientId != tt.clientId {
				t.Errorf("ClientBodyHandler() got = %v, clientId %v", clientId, tt.clientId)
			}
			if secret != tt.secret {
				t.Errorf("ClientBodyHandler() got1 = %v, secret %v", secret, tt.secret)
			}
		})
	}
}
