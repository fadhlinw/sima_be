package config

import (
	"os"
	"sima/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	DSN := os.Getenv("DB_DSN_LOCAL")

	var dberr error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(dberr)
	}
	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.Masjid{})
	DB.AutoMigrate(&models.Inventori{})
	DB.AutoMigrate(models.Transaksi{})
}
