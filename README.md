# OAuth

Modular and simple OAuth implementation for Go.

Authorizing a request
```
req := http.NewRequest("GET", "https://example.com", nil)
c := &Consumer{"key", "secret"}
c.Authorize(req, &Token{"key", "secret"})
```
