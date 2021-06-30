package main

import (
	"gosession/gosession"
	"net/http"
	"time"
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

	http.ListenAndServe(":8080", session.NewGoSession(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			rw.Write([]byte("SHnae"))
		}
	})))
}
