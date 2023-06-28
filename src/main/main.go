package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// corresponds to everything in .env_example
// make sure when updating .env to update here and in func loadEnvVars
type EnvArgs struct {
	DB_URL               string
	SUPABASE_SERVICE_KEY string
}

func main() {
	// start by connecting to db inside of env
	args, err := loadEnvVars()

	if err != nil {
		fmt.Println("Error; unable to load .env file: ", err)
		os.Exit(1)
	}

	fmt.Println(args.DB_URL)
	fmt.Println(args.SUPABASE_SERVICE_KEY)
}

func loadEnvVars() (EnvArgs, error) {
	args := EnvArgs{}

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error; unable to load .env file:", err)
		os.Exit(1)
	}

	args.DB_URL = os.Getenv("DB_URL")
	args.SUPABASE_SERVICE_KEY = os.Getenv("SUPABASE_SERVICE_KEY")

	if args.DB_URL == "" {
		return EnvArgs{}, errors.New("DB_URL not found in .env")
	}

	if args.SUPABASE_SERVICE_KEY == "" {
		return EnvArgs{}, errors.New("SUPABASE_SERVICE_KEY not found in .env")
	}

	return args, nil

}
