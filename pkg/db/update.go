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
