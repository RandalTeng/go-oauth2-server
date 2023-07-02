# OAuth2.0 server for golang.

This repo is base on [github.com/go-oauth2/auth2](https://github.com/go-oauth2/oauth2).

Thanks to `go-oauth2/auth2`'s contributor for their great work.

## This repo for.

### Why new repo not fork the original repo?

I want build an oauth2 server that split rfc definition (which is interfaces usually) and implements.

Everyone can imply each component (access token, code generator, storage component) in their owner implement, 
which is a bit difficult under struct of `go-oauth2/oauth2`.

You have to import the whole package about `go-oauth2/oauth2`, not just the interfaces or access token generator,
there should that user can choose every component which they want. 

### Reason by myself?

Build some amazing work.

## Directories.

1. `definition`: All [RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749) definition is in here.
   This package is necessary for oauth2 server.
2. `errors`: All rfc error description is in here. This package will be autoloaded usually.
3. `generator`: Default access and code string generator. You can implement your own generator.
4. `manager`: Auth-code and AccessToken lifetime manager. Also can be reimplemented.
5. `models`: Default code and token implement. This component will be reimplemented usually.
6. `server`: Default oauth2 server implement, request handler, strongly recommend, NOT reimplement this package.
7. `store`: Default code, token and client id storage, implement for memory saver. This package is one-time-use,
   will be reimplemented on most scenes.
