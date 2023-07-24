package models

import (
	"gorm.io/gorm"
)

type Inventori struct {
	gorm.Model
	NamaBarang      string `json:"nama_barang"`
	Jumlah          int    `json:"jumlah"`
	DeskripsiBarang string `json:"deskripsi_barang"`
	MasjidID        uint   `json:"masjid_id"`
}
