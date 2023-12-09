package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
	"url-shortener/internal/config"
)

type Link struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Alias     string
	Dest      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func New(database config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		database.Host,
		database.User,
		database.Password,
		database.DB,
		database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect Database: %s", err)
	}
	return db
}
