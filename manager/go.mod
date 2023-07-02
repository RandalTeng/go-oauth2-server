module github.com/RandalTeng/go-oauth2-server/manager

go 1.19

require (
	github.com/RandalTeng/go-oauth2-server/definition v1.0.0
	github.com/RandalTeng/go-oauth2-server/errors v1.0.0
	github.com/RandalTeng/go-oauth2-server/generator v1.0.0
	github.com/RandalTeng/go-oauth2-server/models v1.0.0
)

replace (
	github.com/RandalTeng/go-oauth2-server/definition v1.0.0 => ../definition
	github.com/RandalTeng/go-oauth2-server/errors v1.0.0 => ../errors
	github.com/RandalTeng/go-oauth2-server/generator v1.0.0 => ../generator
	github.com/RandalTeng/go-oauth2-server/models v1.0.0 => ../models
)
