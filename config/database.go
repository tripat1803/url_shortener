package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", Env.DATABASE_HOST, Env.DATABASE_USERNAME, Env.DATABASE_PASSWORD, Env.DATABASE_NAME, Env.DATABASE_PORT)
	pgDb := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	var err error
	DB, err = gorm.Open(pgDb, &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting with database. Error:", err)
	}

	log.Println("Connected to database")
}
