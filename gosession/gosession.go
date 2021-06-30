//A go session middleware

package gosession

import (
	"fmt"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId := g.Cookie.Value
		fmt.Println(sessionId)
		//Sets the session
		if _, cookie := r.Cookie(g.Cookie.Name); cookie != nil {
			http.SetCookie(w, NewCookieFromOptions(g.Cookie))
		}
		next.ServeHTTP(w, r)
	})
}
