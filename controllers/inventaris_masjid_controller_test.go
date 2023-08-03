package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sima/config"
	"sima/models"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetInventorisController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/inventoris", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	if assert.NoError(t, GetInventorisController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}
}

func TestGetInventoriController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/inventoris/1", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, GetInventoriController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodGet, "/inventoris/sherlok", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("sherlok")
	error := GetInventoriController(c)

	if assert.Error(t, GetInventoriController(c)) {
		e, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, e.Code)
		assert.Equal(t, "Invalid ID", e.Message)
	}

	request = httptest.NewRequest(http.MethodGet, "/inventoris/100", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := GetInventoriController(c)

	if assert.Error(t, GetInventoriController(c)) {
		e, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, e.Code)
		assert.Equal(t, "Inventori Tidak Ada", e.Message)
	}
}

func TestCreateInventoriController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	book := models.Inventori{
		NamaBarang:      "Sarung Test",
		Jumlah:          9,
		DeskripsiBarang: "Baru",
	}
	inventoriJSON, _ := json.Marshal(book)
	request := httptest.NewRequest(http.MethodPost, "/inventori-create", strings.NewReader(string(inventoriJSON)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	if assert.NoError(t, CreateInventoriController(c)) {
		assert.Equal(t, http.StatusCreated, record.Code)
	}
}

func TestDeleteBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodDelete, "/inventoris/1", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteInventoriController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodDelete, "/inventoris/delete", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("delete")
	error := DeleteInventoriController(c)

	if assert.Error(t, DeleteInventoriController(c)) {
		response, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Invalid ID", response.Message)
	}
}

func TestUpdateBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	book := models.Inventori{
		NamaBarang:      "Sarung Update Test",
		Jumlah:          5,
		DeskripsiBarang: "Baru",
	}
	bookJSON, _ := json.Marshal(book)
	request := httptest.NewRequest(http.MethodPut, "/inventoris/2", strings.NewReader(string(bookJSON)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, UpdateInventoriController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodPut, "/inventoris/update", strings.NewReader(string(bookJSON)))
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("update")
	error := UpdateInventoriController(c)

	if assert.Error(t, UpdateInventoriController(c)) {
		response, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Invalid Request Payload", response.Message)
	}
}
