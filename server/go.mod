module github.com/RandalTeng/oauth2/server

go 1.19

require (
	github.com/RandalTeng/oauth2/definition v1.0.0
	github.com/RandalTeng/oauth2/errors v1.0.0
)

replace (
	github.com/RandalTeng/oauth2/definition v1.0.0 => ../definition
	github.com/RandalTeng/oauth2/errors v1.0.0 => ../errors
)
