//A go session middleware

package gosession

import (
	"encoding/json"
	"net/http"
)

type GosessionMiddleWare struct {
	Cookie              *Cookie
	SessionStoreOptions *SessionStoreOptions
	Store               *Store
}

func Init(c *Cookie, s *SessionStoreOptions) *GosessionMiddleWare {
	// creates the session table if not exists
	store := NewStoreFromOptions(s)
	if store != nil {
		store.CreateSessionTable(s.TableName)
	}
	go RunCleanUp()
	return &GosessionMiddleWare{
		Cookie:              c,
		SessionStoreOptions: s,
		Store:               store,
	}
}

//Returns the middleware
func (g *GosessionMiddleWare) NewGoSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, cookie := r.Cookie(g.Cookie.Name); cookie != nil {
			http.SetCookie(w, NewCookieFromOptions(g.Cookie))
		}
		//ctx := context.WithValue(r.Context(), "data", data)
		//next.ServeHTTP(w, r.WithContext(ctx))
		next.ServeHTTP(w, r)
	})
}

func (g *GosessionMiddleWare) SetValue(val map[string]interface{}) {
	g.Store.CreateSession(g.SessionStoreOptions.TableName, g.Cookie.Value, g.Cookie.Expires, val)
}

func (g *GosessionMiddleWare) GetSessionData() *SessionResponse {
	result := g.Store.GetSessionData(g.Cookie.Value)
	json.Unmarshal([]byte(result.Value), &result.Value)
	return result
}
