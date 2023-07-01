package args

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvArgs struct {
	DB_URL               string
	SUPABASE_SERVICE_KEY string
	PORT                 string
}

type CmdArgs struct {
	Seed bool
}

func LoadEnvVars() (EnvArgs, error) {
	envArgs := EnvArgs{}

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error; unable to load .env file:", err)
		os.Exit(1)
	}

	envArgs.DB_URL = os.Getenv("DB_URL")
	envArgs.SUPABASE_SERVICE_KEY = os.Getenv("SUPABASE_SERVICE_KEY")
	envArgs.PORT = os.Getenv("PORT")

	if envArgs.DB_URL == "" {
		return EnvArgs{}, errors.New("DB_URL not found in .env")
	}

	if envArgs.SUPABASE_SERVICE_KEY == "" {
		return EnvArgs{}, errors.New("SUPABASE_SERVICE_KEY not found in .env")
	}

	return envArgs, nil
}

func LoadCmdArgs() CmdArgs {

	args := CmdArgs{}

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {

			if arg == "--seed" {
				args.Seed = true
			}
		}

	}

	return args
}
