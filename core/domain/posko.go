package domain

import "time"

type Posko struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AlamatPosko   string
	NamaPjPosko   string
	NoTelpPjPosko string
	BencanaID     uint
	Bencana       Bencana
}
