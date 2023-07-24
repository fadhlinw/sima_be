package database

import (
	"sima/config"
	"sima/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Masjid
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (*models.Masjid, error) {
	user := &models.Masjid{}
	if err := config.DB.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserbyEmail(email string) error {
	user := &models.Masjid{}
	if err := config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return err
	}
	return nil
}
func CreateUser(user *models.Masjid) error {

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(user *models.Masjid) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	if err := config.DB.Where("id = ?", id).Delete(&models.Masjid{}).Error; err != nil {
		return err
	}
	return nil
}
func UpdateUser(id int, user *models.Masjid) error {
	if err := config.DB.Model(&models.Masjid{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
