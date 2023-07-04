package middleware

import (
	"errors"
	"log"
	"net/http"

	"com.stimstore/stim-db/src/db"
)

type AuthError struct {
	error
	isAdminError bool
}

func RespondUnauthorized(w http.ResponseWriter, err AuthError) {
	w.WriteHeader(http.StatusUnauthorized)
	log.Println("Error occurred: ", err.Error())

	response := "Unauthorized to access this resource"
	header := http.StatusUnauthorized
	if err.isAdminError {
		response = "Not an admin"
		header = http.StatusForbidden
	}

	w.WriteHeader(header)
	w.Write([]byte(response))
}

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
			RespondUnauthorized(w, AuthError{err, false})
		}

		next.ServeHTTP(w, r)
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := r.Cookie("session")

		err = AuthCore(session, true)

		if err != nil {
			RespondUnauthorized(w, AuthError{err, true})
		}

		next.ServeHTTP(w, r)
	})
}
