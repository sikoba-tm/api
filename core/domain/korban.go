package domain

import (
	"github.com/google/uuid"
	"time"
)

const (
	KONDISI_SEHAT      = "Sehat"
	KONDISI_LUKARINGAN = "Luka Ringan"
	KONDISI_LUKABERAT  = "Luka Berat"
	KONDISI_KRITIS     = "Kritis"
	RENTANG_BALITA     = "Balita"
	RENTANG_ANAK       = "Anak-anak"
	RENTANG_REMAJA     = "Remaja"
	RENTANG_DEWASA     = "Dewasa"
	RENTANG_LANSIA     = "Lansia"
)

type Korban struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Foto           string    `json:"foto"`
	RentangUsia    string    `json:"rentang_usia" form:"rentang_usia"`
	Nama           string    `json:"nama"`
	TempatLahir    string    `json:"tempat_lahir" form:"tempat_lahir"`
	TanggalLahir   time.Time `json:"tanggal_lahir" form:"tanggal_lahir"`
	NamaIbuKandung string    `json:"nama_ibu_kandung" form:"nama_ibu_kandung"`
	Kondisi        string    `json:"kondisi"`
	PoskoID        uint      `json:"-"`
	Posko          Posko
}
