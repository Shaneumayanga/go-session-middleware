# An awful Mysql session store middleware implementation in golang


### Example Usage : 

```go

package main

import (
	"encoding/json"
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
		Expires:  time.Now().Add(time.Hour * 10000),
	}, &gosession.SessionStoreOptions{
		DSN:       "root:@tcp(127.0.0.1:3306)/session?parseTime=true&loc=Local",
		TableName: "sessions",
	})

	router := chi.NewRouter()
	router.Use(session.NewGoSession)
	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello there"))
	})

	router.Get("/set-session", func(rw http.ResponseWriter, r *http.Request) {
		val := make(map[string]interface{})
		val["name"] = "shane"
		session.SetValue(val)
	})

	router.Get("/get-sessionData", func(rw http.ResponseWriter, r *http.Request) {
		session_data := session.GetSessionData()
		response, _ := json.Marshal(session_data)
		rw.Write(response)
	})

	http.ListenAndServe(":8080", router)
}


```

# TODO

- Get sessions
- Delete sessions 
- Update sessions


#### Not intended for production