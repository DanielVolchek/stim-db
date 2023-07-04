package db

import (
	"fmt"
	"log"
)

func GenerateServerAuthToken() {

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
