package db

import (
	"errors"
)

func AuthenticateUserBySession(s string) (*User, error) {
	user := User{}

	err := DB_CONN.Preload("sessions").Where("sessions.token = ?", s).First(&user).Error

	if err != nil {
		return nil, errors.Join(errors.New("Provided token is unauthorized"), err)
	}

	return &user, nil
}

func ValidateEmailToken(t string) error {

	if t == "" {
		return errors.New("Token required")
	}

	return DB_CONN.First(&EmailToken{Token: t}).Error
}
