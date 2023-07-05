package routes

import (
	"fmt"
	"net/http"

	"github.com/danielvolchek/stim-db/pkg/db"
)

var EmailRoute Route = Route{
	route:        "/auth",
	finalHandler: EmailRouteHandler,
	middleware:   nil,
}

var EmailRouteHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetHandler(w, r)
	} else if r.Method == "POST" {
		PostHandler(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("method %s not allowed", r.Method)))
	}
})

// checks to see if a token exists
func GetHandler(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("token")

	err := db.ValidateEmailToken(token)

	access := "A"

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		access = "una"
	}

	// Auth vs unauth string builder
	w.Write([]byte(fmt.Sprintf("%suthorized to access this resource", access)))
}

// adds a new token
func PostHandler(w http.ResponseWriter, r *http.Request) {

}
