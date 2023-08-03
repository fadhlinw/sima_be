package controllers

import (
	"fmt"
	"net/http"
	"sima/config"
	"sima/database"
	"sima/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetTransaksisController(c echo.Context) error {

	invs, err := database.GetTransaksis()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"transaksis": invs,
	})
}

func GetTransaksiController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	inv, err := database.GetTransaksi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Inventori Tidak Ada")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"transaksis": inv,
	})
}

func GetLastTransaksiIDByMasjidID(idmasjid int) (uint, error) {
	var maxID uint
	err := config.DB.Model(&models.Transaksi{}).Where("masjid_id = ?", idmasjid).Order("id desc").Select("id").Limit(1).Scan(&maxID).Error
	if err != nil {
		return 0, err
	}

	return maxID, nil
}

func GetLastTotalKasByID(id int) (float64, error) {
	var lastTotalKas float64
	err := config.DB.Model(&models.Transaksi{}).Where("id = ?", id).Order("id desc").Limit(1).Pluck("total_kas", &lastTotalKas).Error
	if err != nil {
		return 0, err
	}

	return lastTotalKas, nil
}

func CreateTransaksiController(c echo.Context) error {
	var transaksi models.Transaksi

	idmasjid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	fileHeader, err := c.FormFile("photo")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get the file from request")
	}

	uploadedURL, err := uploadFile(fileHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upload file to S3")
	}

	if err := c.Bind(&transaksi); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	transaksi.MasjidID = uint(idmasjid)

	transaksi.PhotoURL = uploadedURL

	// Mendapatkan nilai TotalKas sebelum menyimpan transaksi ke database
	lastTransaksiID, err := GetLastTransaksiIDByMasjidID(idmasjid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get last transaksi ID")
	}

	lastTotalKas, err := GetLastTotalKasByID(int(lastTransaksiID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get last total kas")
	}

	transaksiResponse := models.Transaksi{
		NamaTransaksi:      transaksi.NamaTransaksi,
		Tanggal:            transaksi.Tanggal,
		JumlahTransaksi:    transaksi.JumlahTransaksi,
		DeskripsiTransaksi: transaksi.DeskripsiTransaksi,
		TotalKas:           transaksi.TotalKas,
		PhotoURL:           transaksi.PhotoURL,
	}

	// Log nilai sebelum perhitungan
	fmt.Println("Before calculation:")
	fmt.Println("lastTotalKas:", lastTotalKas)
	fmt.Println("JenisTransaksi:", transaksi.JenisTransaksi)
	fmt.Println("JumlahTransaksi:", transaksi.JumlahTransaksi)

	if transaksi.JenisTransaksi == "Masuk" {
		transaksi.JumlahTransaksi = +transaksi.JumlahTransaksi
	} else if transaksi.JenisTransaksi == "Keluar" {
		transaksi.JumlahTransaksi = -transaksi.JumlahTransaksi
	}

	// Log nilai setelah perhitungan
	fmt.Println("After calculation:")
	fmt.Println("lastTotalKas:", lastTotalKas)
	fmt.Println("JenisTransaksi:", transaksi.JenisTransaksi)
	fmt.Println("JumlahTransaksi:", transaksi.JumlahTransaksi)

	transaksi.TotalKas = lastTotalKas + transaksi.JumlahTransaksi

	err = database.CreateTransaksi(&transaksi)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":     "success",
		"transaksis": transaksiResponse,
	})
}

func DeleteTransaksiController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	err = database.DeleteTransaksi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Inventori Tidak Ada")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Inventori Deleted",
	})
}

func UpdateTransaksiController(c echo.Context) error {
	var transaksi models.Transaksi

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if err := c.Bind(&transaksi); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Payload")
	}

	err = database.UpdateTransaksi(id, &transaksi)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Inventori Updated",
	})
}
