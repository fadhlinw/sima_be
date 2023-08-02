package database

import (
	"sima/config"
	"sima/models"
)

func GetTransaksis() (interface{}, error) {
	var transaksis []models.Transaksi
	if err := config.DB.Find(&transaksis).Error; err != nil {
		return nil, err
	}
	return transaksis, nil
}

func GetTransaksi(id int) (*models.Transaksi, error) {
	transaksi := &models.Transaksi{}
	if err := config.DB.Where("id = ?", id).Find(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}

func CreateTransaksi(tr *models.Transaksi) error {
	if err := config.DB.Create(&tr).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTransaksi(id int) error {
	transaksi := &models.Inventori{}
	if err := config.DB.Where("id = ?", id).Delete(transaksi).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTransaksi(id int, tr *models.Transaksi) error {
	if err := config.DB.Model(&models.Transaksi{}).Where("id = ?", id).Updates(tr).Error; err != nil {
		return err
	}
	return nil
}
