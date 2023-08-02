package database

import (
	"sima/config"
	"sima/models"
)

func GetInventoris() (interface{}, error) {
	var inventoris []models.Inventori
	if err := config.DB.Find(&inventoris).Error; err != nil {
		return nil, err
	}
	return inventoris, nil
}

func GetInventori(id int) (*models.Inventori, error) {
	inventori := &models.Inventori{}
	if err := config.DB.Where("id = ?", id).First(inventori).Error; err != nil {
		return nil, err
	}
	return inventori, nil
}

func CreateInventori(inv *models.Inventori) error {
	if err := config.DB.Create(&inv).Error; err != nil {
		return err
	}
	return nil
}

func DeleteInventori(id int) error {
	inventori := &models.Inventori{}
	if err := config.DB.Where("id = ?", id).Delete(inventori).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInventori(id int, inv *models.Inventori) error {
	if err := config.DB.Model(&models.Inventori{}).Where("id = ?", id).Updates(inv).Error; err != nil {
		return err
	}
	return nil
}
