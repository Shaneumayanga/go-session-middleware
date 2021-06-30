package main

import (
	"gosession/gosession"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {

	session := gosession.Init(&gosession.Cookie{
		Name:     "session",
		Value:    "thisisthesessionId4",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Second * 60),
	}, &gosession.SessionStoreOptions{})

	router := chi.NewRouter()
	router.Use(session.NewGoSession)
	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello there!"))
	})
	router.Get("/session", func(rw http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("session")
		rw.Write([]byte(cookie.Value))
	})
	http.ListenAndServe(":8080", router)
}
