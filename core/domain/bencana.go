package domain

import (
	"time"
)

type Bencana struct {
	ID              uint `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	NamaBencana     string    `json:"nama_bencana"`
	Lokasi          string    `json:"lokasi"`
	TanggalKejadian time.Time `json:"tanggal_bencana"`
}
