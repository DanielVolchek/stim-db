package routes

import (
	"net/http"

	"com.stimstore/stim-db/src/router/middleware"
)

var userHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("good job my friend it is nice to meet you!"))
})

var UserRoute = Route{"/user", userHandler, []Middleware{
	middleware.AuthMiddleware,
	middleware.AdminMiddleware,
}}
