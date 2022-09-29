package domain

import "time"

type Petugas struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	username       string
	hashedPassword string
}
