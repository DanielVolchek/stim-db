package main

import (
	// std

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
	envArgs := args.EnvArgs
	cmdArgs := args.CmdArgs

	if cmdArgs.Migrate {
		db.MigrateDB()
		return
	}

	router.StartHttpClient(envArgs.PORT)

}
