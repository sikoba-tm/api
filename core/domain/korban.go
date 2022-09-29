package domain

import "time"

const (
	SEHAT      = "Sehat"
	LUKARINGAN = "Luka Ringan"
	LUKABERAT  = "Luka Berat"
	KRITIS     = "Kritis"
)

type Korban struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Nama           string
	PoB            string
	DoB            time.Time
	NamaIbuKandung string
	Foto           string
	Kondisi        string
	PoskoID        uint
	Posko          Posko
}
