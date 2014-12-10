# OAuth

Modular and simple OAuth implementation for Go.

## Installation

`go get github.com/nhjk/oauth`

## Documentation

http://godoc.org/github.com/nhjk/oauth

## Authorizing a request

```go
req := http.NewRequest("GET", "https://example.com", nil)
c := &oauth.Consumer{*ckey*, *csecret*}
c.Authorize(req, &oauth.Token{*tkey*, *tsecret*})
```
