module github.com/RandalTeng/oauth2/store

go 1.19

require (
	github.com/RandalTeng/go-oauth2-server/definition v1.0.0
)

replace (
	github.com/RandalTeng/go-oauth2-server/definition v1.0.0 => ../definition
)
