package domain

import "time"

type Posko struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	NamaPj    string `json:"nama_pj"`
	NoTelpPj  string `json:"notelp_pj"`
	BencanaID uint
	Bencana   Bencana `json:"-"`
}
