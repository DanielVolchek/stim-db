package routes

import (
	"fmt"
	"net/http"

	"github.com/danielvolchek/stim-db/pkg/router/middleware"
)

type Middleware func(http.Handler) http.Handler

type Route struct {
	route        string
	finalHandler http.Handler
	// Middlewares are called in reverse order
	// (index[0](index[1](finalHandler)))
	middleware []Middleware
}

func (route *Route) ConstructRouteHandler(handler *http.ServeMux) {
	if len(route.middleware) > 0 {
		var handleFunc http.Handler = route.finalHandler

		// Apply middlewares in reverse order
		for i := 0; i < len(route.middleware); i++ {
			handleFunc = route.middleware[i](handleFunc)
		}

		handler.Handle(route.route, handleFunc)
	} else {
		handler.Handle(route.route, route.finalHandler)
	}
}

var IndexFinal = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit index route")
	w.Write([]byte("Hello from the index"))
})

var IndexRoute Route = Route{
	"/", IndexFinal, []Middleware{middleware.AuthMiddleware},
}

func Router(handler *http.ServeMux) {
	// val := middleware.AuthMiddleware(IndexRoute)
	// handler.Handle("/")
	// handler.Handle(UserRoute.route, UserRoute.handler)
	IndexRoute.ConstructRouteHandler(handler)
	UserRoute.ConstructRouteHandler(handler)
}

func handlerFunc(res http.ResponseWriter, req *http.Request) {
	// extendedRes := ExtendedResponseWriter{ResponseWriter: res}
}
