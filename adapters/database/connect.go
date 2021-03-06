package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() Database {
	dsn, ok := os.LookupEnv("POSTGRES_CONNECTION")
	if !ok {
		log.Fatalln("POSTGRES_CONNECTION does not exist in environment variables!")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		log.Fatalln("Failed to connect to database!")
	}

	return New(db)
}
