package controllers

import (
	"net/http"
	"sima/database"
	"sima/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetInventorisController(c echo.Context) error {

	invs, err := database.GetInventoris()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  invs,
	})
}

func GetInventoriController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	inv, err := database.GetInventori(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Inventori Tidak Ada")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   inv,
	})
}

func CreateInventoriController(c echo.Context) error {
	var inventori models.Inventori

	idmasjid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if err := c.Bind(&inventori); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	inventori.MasjidID = uint(idmasjid)

	inventoriResponse := models.Inventori{
		NamaBarang:      inventori.NamaBarang,
		Jumlah:          inventori.Jumlah,
		DeskripsiBarang: inventori.DeskripsiBarang,
	}

	err = database.CreateInventori(&inventori)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   inventoriResponse,
	})
}

func DeleteInventoriController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	err = database.DeleteInventori(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Inventori Tidak Ada")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Inventori Deleted",
	})
}

func UpdateInventoriController(c echo.Context) error {
	var inventori models.Inventori

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if err := c.Bind(&inventori); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Payload")
	}

	err = database.UpdateInventori(id, &inventori)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Inventori Updated",
	})
}
