package middleware

import (
	"errors"
	"log"
	"net/http"

	"com.stimstore/stim-db/src/db"
)

func AuthCore(session *http.Cookie, needsAdmin bool) error {

	err := session.Valid()

	if err != nil {
		return err
	}

	user, err := db.AuthenticateUserBySession(session.String())

	if err != nil {

		return err
	}

	if needsAdmin && user.Role != "ADMIN" {
		return errors.New("User is unauthorized as an admin")
	}

	return nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := r.Cookie("session")

		err = AuthCore(session, false)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Error occurred: ", err.Error())
			w.Write([]byte("Unauthorized to access this resource given information provided"))
		}

		next.ServeHTTP(w, r)
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := r.Cookie("session")

		err = AuthCore(session, true)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Error occurred: ", err.Error())
			w.Write([]byte("Unauthorized to access this resource given information provided"))
		}

		next.ServeHTTP(w, r)
	})
}
