//A go session middleware

package gosession

import (
	"net/http"
)

type GosessionMiddleWare struct {
	Cookie              *Cookie
	SessionStoreOptions *SessionStoreOptions
}

func Init(c *Cookie, s *SessionStoreOptions) *GosessionMiddleWare {
	return &GosessionMiddleWare{
		Cookie:              c,
		SessionStoreOptions: s,
	}
}

func (g *GosessionMiddleWare) NewGoSession(next http.Handler) http.Handler {
	// creates the session table if not exists
	store := NewStoreFromOptions(g.SessionStoreOptions)
	store.CreateSessionTable(g.SessionStoreOptions.TableName)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, cookie := r.Cookie(g.Cookie.Name); cookie != nil {
			http.SetCookie(w, NewCookieFromOptions(g.Cookie))
		}
		next.ServeHTTP(w, r)
	})
}
