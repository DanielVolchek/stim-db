package db

import "log"

func ValidateServerToken(token any) bool {
	var err error

	switch t := token.(type) {
	case string:
		err = DB_CONN.First(&ServerAuthToken{Token: t}).Error
	case ServerAuthToken:
		err = DB_CONN.First(&t).Error
	case *ServerAuthToken:
		err = DB_CONN.First(t).Error
	default:
		log.Fatal("typeof token must be string or ServerAuthToken")
	}

	if err != nil {
		log.Fatal(err)
	}

	if DB_CONN.RowsAffected == 0 {
		return false
	}

	return true
}
