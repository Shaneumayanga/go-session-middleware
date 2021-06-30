package gosession

import (
	"net/http"
	"time"
)

type Cookie struct {
	//Session cookie Name
	Name string
	//Session id
	Value    string
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
	Expires  time.Time
}

func NewCookieFromOptions(c *Cookie) *http.Cookie {
	return &http.Cookie{
		Name:     c.Name,
		Value:    c.Value,
		Path:     c.Path,
		Domain:   c.Domain,
		Secure:   c.Secure,
		HttpOnly: c.HttpOnly,
		Expires:  c.Expires,
	}
}
