# OAuth 1.0a

Modular and simple OAuth 1.0a implementation for Go.

## Installation

`go get github.com/nhjk/oauth`

## Documentation

http://godoc.org/github.com/nhjk/oauth

## Authorizing a request

For a more detailed example, see [examples/twitter.go](https://github.com/nhjk/oauth/blob/master/examples/twitter.go).

```go
req := http.NewRequest("GET", "https://example.com", nil)
c := &oauth.Consumer{*ckey*, *csecret*}
c.Authorize(req, &oauth.Token{*tkey*, *tsecret*})
```
