package config

import "os"

func CheckEnvLoaded() bool {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USERNAME") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_ADDRESS") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_SCHEMA") == "" {
		return false
	}

	return true
}

var (
	SERVER_ADDRESS = os.Getenv("SERVER_ADDRESS")
	SERVER_PORT    = os.Getenv("SERVER_PORT")
	DB_USERNAME    = os.Getenv("DB_USERNAME")
	DB_PASSWORD    = os.Getenv("DB_PASSWORD")
	DB_ADDRESS     = os.Getenv("DB_ADDRESS")
	DB_PORT        = os.Getenv("DB_PORT")
	DB_SCHEMA      = os.Getenv("DB_SCHEMA")
)
