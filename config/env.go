package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvData struct {
	PORT              string
	DATABASE_HOST     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_PORT     string
	DATABASE_NAME     string
	JWT_SECRET        string
}

var Env EnvData

func GetConfigEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env")
	}

	Env.PORT = os.Getenv("PORT")
	Env.DATABASE_HOST = os.Getenv("DATABASE_HOST")
	Env.DATABASE_USERNAME = os.Getenv("DATABASE_USERNAME")
	Env.DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	Env.DATABASE_NAME = os.Getenv("DATABASE_NAME")
	Env.DATABASE_PORT = os.Getenv("DATABASE_PORT")
	Env.JWT_SECRET = os.Getenv("JWT_SECRET")
}
