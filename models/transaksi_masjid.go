package models

import (
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	NamaTransaksi      string  `json:"nama_transaksi"`
	JenisTransaksi     string  `json:"jenis_transaksi"`
	Tanggal            string  `json:"tanggal"`
	JumlahTransaksi    float64 `json:"jumlah_transaksi"`
	DeskripsiTransaksi string  `json:"deskripsi_transaksi"`
	TotalKas           float64 `json:"total_kas"`
	MasjidID           uint    `json:"masjid_id"`
	PhotoURL           string  `json:"photo_url"`
}
