package server

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/RandalTeng/go-oauth2-server/definition"
)

func TestServer_GetAccessToken(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		ctx context.Context
		gt  definition.GrantType
		tgr *definition.TokenGenerateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    definition.TokenInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			got, err := s.GetAccessToken(tt.args.ctx, tt.args.gt, tt.args.tgr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetAuthorizeData(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		rt definition.ResponseType
		ti any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			if got := s.GetAuthorizeData(tt.args.rt, tt.args.ti); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthorizeData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetAuthorizeToken(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		ctx context.Context
		req *AuthorizeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			got, err := s.GetAuthorizeToken(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthorizeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthorizeToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetErrorData(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
		want1  int
		want2  http.Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			got, got1, got2 := s.GetErrorData(tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetErrorData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetErrorData() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetErrorData() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestServer_GetRedirectURI(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		req  *AuthorizeRequest
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			got, err := s.GetRedirectURI(tt.args.req, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRedirectURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRedirectURI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetTokenData(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		ti definition.TokenInfo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			if got := s.GetTokenData(tt.args.ti); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTokenData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_HandleAuthorizeRequest(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			if err := s.HandleAuthorizeRequest(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandleAuthorizeRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_HandleTokenRequest(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			if err := s.HandleTokenRequest(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandleTokenRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_ValidationAuthorizeRequest(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AuthorizeRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			got, err := s.ValidationAuthorizeRequest(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidationAuthorizeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidationAuthorizeRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ValidationBearerToken(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    definition.TokenInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			got, err := s.ValidationBearerToken(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidationBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidationBearerToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ValidationTokenRequest(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    definition.GrantType
		want1   *definition.TokenGenerateRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			got, got1, err := s.ValidationTokenRequest(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidationTokenRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidationTokenRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidationTokenRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestServer_token(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		w          http.ResponseWriter
		data       map[string]interface{}
		header     http.Header
		statusCode []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			if err := s.token(tt.args.w, tt.args.data, tt.args.header, tt.args.statusCode...); (err != nil) != tt.wantErr {
				t.Errorf("token() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_tokenError(t *testing.T) {
	type fields struct {
		Config                       *Config
		Manager                      definition.Manager
		ClientInfoHandler            ClientInfoHandler
		ClientAuthorizedHandler      ClientAuthorizedHandler
		ClientScopeHandler           ClientScopeHandler
		UserAuthorizationHandler     UserAuthorizationHandler
		PasswordAuthorizationHandler PasswordAuthorizationHandler
		RefreshingValidationHandler  RefreshingValidationHandler
		PreRedirectErrorHandler      PreRedirectErrorHandler
		RefreshingScopeHandler       RefreshingScopeHandler
		ResponseErrorHandler         ResponseErrorHandler
		InternalErrorHandler         InternalErrorHandler
		ExtensionFieldsHandler       ExtensionFieldsHandler
		AuthorizeScopeHandler        AuthorizeScopeHandler
		ResponseTokenHandler         ResponseTokenHandler
	}
	type args struct {
		w   http.ResponseWriter
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Config:                       tt.fields.Config,
				Manager:                      tt.fields.Manager,
				ClientInfoHandler:            tt.fields.ClientInfoHandler,
				ClientAuthorizedHandler:      tt.fields.ClientAuthorizedHandler,
				ClientScopeHandler:           tt.fields.ClientScopeHandler,
				UserAuthorizationHandler:     tt.fields.UserAuthorizationHandler,
				PasswordAuthorizationHandler: tt.fields.PasswordAuthorizationHandler,
				RefreshingValidationHandler:  tt.fields.RefreshingValidationHandler,
				PreRedirectErrorHandler:      tt.fields.PreRedirectErrorHandler,
				RefreshingScopeHandler:       tt.fields.RefreshingScopeHandler,
				ResponseErrorHandler:         tt.fields.ResponseErrorHandler,
				InternalErrorHandler:         tt.fields.InternalErrorHandler,
				ExtensionFieldsHandler:       tt.fields.ExtensionFieldsHandler,
				AuthorizeScopeHandler:        tt.fields.AuthorizeScopeHandler,
				ResponseTokenHandler:         tt.fields.ResponseTokenHandler,
			}
			if err := s.tokenError(tt.args.w, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("tokenError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
