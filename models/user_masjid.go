package models

import (
	"gorm.io/gorm"
)

type Masjid struct {
	gorm.Model
	NamaMasjid        string `json:"nama_masjid"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	NamaTakmir        string `json:"nama_takmir"`
	AlamatMasjid      string `json:"alamat_masjid"`
	KontakPerson      string `json:"kontak_person"`
	ProfilURL         string `json:"profil_url"`
	Inventaris        []Inventori
	TransaksiKeuangan []Transaksi
}

type MasjidResponse struct {
	ID           int    `json:"id"`
	NamaMasjid   string `json:"nama_masjid"`
	Email        string `json:"email"`
	NamaTakmir   string `json:"nama_takmir"`
	AlamatMasjid string `json:"alamat_masjid"`
	KontakPerson string `json:"kontak_person"`
	ProfilURL    string `json:"profil_url"`
	Token        string `json:"token"`
}
