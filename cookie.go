package sessions

import (
	"github.com/gorilla/sessions"
)

var _ sessions.Store = (*cookieStore)(nil)

func NewCookieStore(keyPairs ...[]byte) Store {
	return &cookieStore{sessions.NewCookieStore(keyPairs...)}
}

type cookieStore struct {
	*sessions.CookieStore
}

func (c *cookieStore) Options(options AweOptions) {
	c.CookieStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
