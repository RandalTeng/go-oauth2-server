package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RandalTeng/go-oauth2-server/definition"
	"github.com/RandalTeng/go-oauth2-server/errors"
)

type (
	// ClientInfoHandler get client info from request
	ClientInfoHandler func(r *http.Request) (clientID, clientSecret string, err error)

	// ClientAuthorizedHandler check the client allows to use this authorization grant type
	ClientAuthorizedHandler func(clientID string, grant definition.GrantType) (allowed bool, err error)

	// ClientScopeHandler check the client allows to use scope
	ClientScopeHandler func(tgr *definition.TokenGenerateRequest) (allowed bool, err error)

	// UserAuthorizationHandler get user id from request authorization
	UserAuthorizationHandler func(w http.ResponseWriter, r *http.Request) (userID string, err error)

	// PasswordAuthorizationHandler get user id from username and password
	PasswordAuthorizationHandler func(ctx context.Context, clientID, username, password string) (userID string, err error)

	// RefreshingScopeHandler check the scope of the refreshing token
	RefreshingScopeHandler func(tgr *definition.TokenGenerateRequest, oldScope string) (allowed bool, err error)

	// RefreshingValidationHandler check if refresh_token is still valid. eg no revocation or other
	RefreshingValidationHandler func(ti definition.TokenInfo) (allowed bool, err error)

	// ResponseErrorHandler response error handing
	ResponseErrorHandler func(re *errors.Response)

	// InternalErrorHandler internal error handing
	InternalErrorHandler func(err error) (re *errors.Response)

	// PreRedirectErrorHandler is used to override "redirect-on-error" behavior
	PreRedirectErrorHandler func(w http.ResponseWriter, req *AuthorizeRequest, err error) error

	// AuthorizeScopeHandler set the authorized scope
	AuthorizeScopeHandler func(w http.ResponseWriter, r *http.Request) (scope string, err error)

	// ExtensionFieldsHandler in response to the access token with the extension of the field
	ExtensionFieldsHandler func(ti definition.TokenInfo) (fieldsValue map[string]interface{})

	// ResponseTokenHandler response token handing
	ResponseTokenHandler func(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error
)

// ClientBodyHandler get client data from request body.
func ClientBodyHandler(r *http.Request) (string, string, error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var decoder *json.Decoder
		if r.GetBody != nil {
			bd, _ := r.GetBody()
			decoder = json.NewDecoder(bd)
		} else {
			decoder = json.NewDecoder(r.Body)
		}
		client := struct {
			ClientId     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
		}{}
		_ = decoder.Decode(&client)
		return client.ClientId, client.ClientSecret, nil
	default:
		clientID := r.Form.Get("client_id")
		if clientID == "" {
			return "", "", errors.ErrInvalidClient
		}
		clientSecret := r.Form.Get("client_secret")
		return clientID, clientSecret, nil
	}
}

// ClientBasicHandler get client data from basic authorization
func ClientBasicHandler(r *http.Request) (string, string, error) {
	username, password, ok := r.BasicAuth()
	if !ok {
		return "", "", errors.ErrInvalidClient
	}
	return username, password, nil
}
