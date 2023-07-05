package routes

import (
	"fmt"
	"net/http"

	"github.com/danielvolchek/stim-db/pkg/db"
	"github.com/danielvolchek/stim-db/pkg/router/middleware"
)

var TokenRoute Route = Route{
	route:        "/admin/token",
	finalHandler: http.HandlerFunc(TokenFinalHandler),
	middleware:   []Middleware{middleware.AdminMiddleware},
}

func TokenGetHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if !db.ValidateServerToken(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Requires server token to be provided, generate on server"))
	}
}

// func TokenPostHandler(w http.ResponseWriter, r *http.Request) {
//
// }

func TokenFinalHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		TokenGetHandler(w, r)
	// case "POST":
	// 	TokenPostHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("Method %s not allowed", r.Method)))
	}
}
