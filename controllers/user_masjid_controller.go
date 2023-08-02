package controllers

import (
	"net/http"
	"sima/database"
	"sima/middleware"
	"sima/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	user, err := database.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Tidak Ada")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func LoginUserController(c echo.Context) error {
	var user models.Masjid
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := database.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := middleware.CreateToken(int(user.ID), user.NamaMasjid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	UserResponse := models.MasjidResponse{ID: int(user.ID), NamaMasjid: user.NamaMasjid, Email: user.Email, NamaTakmir: user.NamaTakmir, AlamatMasjid: user.AlamatMasjid, KontakPerson: user.KontakPerson, Token: token}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   UserResponse,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.Masjid

	fileHeader, err := c.FormFile("photo")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get the file from request")
	}

	uploadedURL, err := uploadFile(fileHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upload file to S3")
	}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	user.ProfilURL = uploadedURL

	err = database.GetUserbyEmail(user.Email)
	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Email Sudah Terdaftar")
	}

	masjidResponse := models.MasjidResponse{
		ID:           int(user.ID),
		NamaMasjid:   user.NamaMasjid,
		Email:        user.Email,
		NamaTakmir:   user.NamaTakmir,
		AlamatMasjid: user.AlamatMasjid,
		KontakPerson: user.KontakPerson,
		ProfilURL:    user.ProfilURL,
	}

	err = database.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   masjidResponse,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	err = database.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Tidak Ada")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "User Deleted",
	})
}

func UpdateUserController(c echo.Context) error {
	var user models.Masjid
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Payload")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	err = database.UpdateUser(id, &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "User Updated",
	})
}
