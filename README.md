# OAuth2.0 server for golang.

This repo is base on [github.com/go-oauth2/auth2](https://github.com/go-oauth2/oauth2).

Thanks to `go-oauth2/auth2`'s contributor for great work.

## This repo for.

### Why new repo not fork the original repo?

I want build an oauth2 server that split rfc definition (which is interfaces usually) and implements.

Everyone can imply each component (access token, code generator, storage component) in their owner implement, 
which is a bit difficult under struct of `go-oauth2/oauth2`.

You have to import the whole package about `go-oauth2/oauth2`, not just the interfaces or access token generator,
there should that user can choose every component which they want. 

### Reason by myself?

Build some amazing work.
