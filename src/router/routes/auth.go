package routes

import (
	"net/http"

	"com.stimstore/stim-db/src/router/middleware"
)

var AuthRoute Route = Route{
	route:        "/auth",
	finalHandler: AuthRouteHandler,
	middleware:   []Middleware{middleware.AuthMiddleware},
}

var AuthRouteHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
})
