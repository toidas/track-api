package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return os.Getenv(name)
}
