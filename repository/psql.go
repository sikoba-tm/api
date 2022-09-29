package repository

import (
	"github.com/sikoba-tm/api/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(dbURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to PostgreSQL Database")

	_ = db.AutoMigrate(
		&domain.Petugas{},
		&domain.Bencana{},
		&domain.Posko{},
		&domain.Korban{},
	)
	log.Println("Model migration successful")

	return db
}
