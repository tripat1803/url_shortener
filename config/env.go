package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT              string
	DATABASE_HOST     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_PORT     string
	DATABASE_NAME     string
}

var env Env

func GetConfigEnv() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env")
	}

	env.PORT = os.Getenv("PORT")
	env.DATABASE_HOST = os.Getenv("DATABASE_HOST")
	env.DATABASE_USERNAME = os.Getenv("DATABASE_USERNAME")
	env.DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	env.DATABASE_NAME = os.Getenv("DATABASE_NAME")
	env.DATABASE_PORT = os.Getenv("DATABASE_PORT")

	return env
}
