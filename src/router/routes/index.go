package routes

import (
	"fmt"
	"net/http"

	"com.stimstore/stim-db/src/router/middleware"
)

var IndexRoute = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit index route")
	w.Write([]byte("Hello from the index"))
})

func Router(handler *http.ServeMux) {
	handler.Handle("/", middleware.AuthMiddleware(IndexRoute))
}

func handlerFunc(res http.ResponseWriter, req *http.Request) {
	// extendedRes := ExtendedResponseWriter{ResponseWriter: res}
}
