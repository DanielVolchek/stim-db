package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("cookies are ", r.Cookies())

		_, err := r.Cookie("session")

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Requires session cookie"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
