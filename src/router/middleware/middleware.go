package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in session")
		// fmt.Println("cookies are ", r.Cookies())

		// _, err := r.Cookie("session")
		//
		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte("Requires session cookie"))
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in admin")

		// _, err := r.Cookie("admin")

		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte("Requires session cookie"))
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
