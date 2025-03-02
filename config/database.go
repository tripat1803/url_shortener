package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	env := GetConfigEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.DATABASE_HOST, env.DATABASE_USERNAME, env.DATABASE_PASSWORD, env.DATABASE_NAME, env.DATABASE_PORT)
	pgDb := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(pgDb, &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting with database. Error:", err)
	}

	log.Println("Connected to database")
	return db
}
