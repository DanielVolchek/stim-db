package main

import (
  // std
	"errors"
	"fmt"
	"os"

  // internal
  "com.stimstore/stim-db/src/db"

  // external
	"github.com/joho/godotenv"
)

// corresponds to everything in .env_example
// make sure when updating .env to update here and in func loadEnvVars


func main() {
	// start by connecting to db inside of env
	envArgs, err := loadEnvVars()

	if err != nil {
		fmt.Println("Error; unable to load .env file: ", err)
		os.Exit(1)
	}

  cmdArgs := checkArgs()

  fmt.Println(envArgs)
  fmt.Println(cmdArgs)

  if (cmdArgs.seed){
    db.SeedDB(envArgs)
    os.Exit(0)
  }

}

