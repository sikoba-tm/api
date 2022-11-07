package domain

import "time"

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
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Foto           string    `json:"foto"`
	RentangUsia    string    `json:"rentang_usia"`
	Nama           string    `json:"nama"`
	TempatLahir    string    `json:"tempat_lahir"`
	TanggalLahir   time.Time `json:"tangal_lahir"`
	NamaIbuKandung string    `json:"nama_ibu_kandung"`
	Kondisi        string    `json:"kondisi"`
	PoskoID        uint
	Posko          Posko
}
