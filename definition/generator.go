package definition

import (
	"context"
	"net/http"
	"time"
)

type (
	// GenerateBasic provide the basis of the generated token data
	GenerateBasic struct {
		Client    ClientInfo
		UserID    string
		Scope     string
		CreatedAt time.Time
		ExpiredAt time.Time
		Request   *http.Request
	}

	// GenerateTokenBody generator's result, if there is a AccessIdentifier (which is set for jwt,
	// the result token is too large), you can just save the identifier to storage
	GenerateTokenBody struct {
		Access           string
		AccessIdentifier string
		Refresh          string
	}

	// AuthorizeGenerate generate the authorization code interface
	AuthorizeGenerate interface {
		Token(ctx context.Context, data *GenerateBasic) (code string, err error)
	}

	// AccessGenerate generate the access and refresh tokens interface
	AccessGenerate interface {
		Token(ctx context.Context, data *GenerateBasic, isGenRefresh bool) (body GenerateTokenBody, err error)
	}
)
