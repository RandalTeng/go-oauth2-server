module github.com/RandalTeng/go-oauth2-server/server

go 1.19

require (
	github.com/RandalTeng/go-oauth2-server/definition v1.0.0
	github.com/RandalTeng/go-oauth2-server/errors v1.0.0
)

replace (
	github.com/RandalTeng/go-oauth2-server/definition v1.0.0 => ../definition
	github.com/RandalTeng/go-oauth2-server/errors v1.0.0 => ../errors
)
