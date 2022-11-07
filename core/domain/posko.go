package domain

import "time"

type Posko struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	NamaPj    string `json:"penanggung_jawab"`
	NoTelpPj  string `json:"no_telp"`
	BencanaID uint
	Bencana   Bencana
}
