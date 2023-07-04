package db

import (
	"errors"
	"fmt"
	"log"

	"com.stimstore/stim-db/src/args"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_CONN = ConnectDB(args.EnvArgs)

func ConnectDB(argList args.EnvArgsType) *gorm.DB {
	db, err := gorm.Open(postgres.Open(argList.DB_URL), &gorm.Config{})

	if err != nil {
		log.Print("Failed to instantiate DB: ")
		log.Fatal(err)
	}

	return db
}

func Find(t *interface{}) (tx *gorm.DB) {
	return DB_CONN.Find(t)
}

func AuthenticateUserBySession(s string) (*User, error) {
	user := User{}

	err := DB_CONN.Preload("sessions").Where("sessions.token = ?", s).First(&user).Error

	if err != nil {
		return nil, errors.Join(errors.New("Provided token is unauthorized"), err)
	}

	return &user, nil
}

func CreateNewUser() {
	user := User{Username: "dn", Role: "USER", PasswordHash: "hi"}
	err := DB_CONN.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
}

func CreateNewSessionOnUser() error {

	user := &User{}
	DB_CONN.Where(&User{Username: "daniel"}).First(&user)

	fmt.Printf("%+v\n", user)

	return nil
}
