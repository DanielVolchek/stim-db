package main

import (
	// std

	"fmt"
	"os"

	// internal
	"com.stimstore/stim-db/src/args"
	"com.stimstore/stim-db/src/db"
	"com.stimstore/stim-db/src/router"
	// external
)

// corresponds to everything in .env_example
// make sure when updating .env to update here and in func loadEnvVars

func main() {
	// start by connecting to db inside of env
	envArgs, err := args.LoadEnvVars()

	if err != nil {
		fmt.Println("Error; unable to load .env file: ", err)
		os.Exit(1)
	}

	cmdArgs := args.LoadCmdArgs()

	if cmdArgs.Seed {
		err := db.SeedDB(envArgs)
		if err != nil {
			fmt.Println("Error during seed: ", err)
			fmt.Println("Exiting")
			os.Exit(1)
		}
	}

	fmt.Println(envArgs)
	fmt.Println(cmdArgs)

	router.StartHttpClient(envArgs.PORT)

}
