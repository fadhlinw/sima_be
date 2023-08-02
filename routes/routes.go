package routes

import (
	"os"
	"sima/controllers"
	"sima/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middleware.LogMiddleware(e)

	e.GET("/masjids", controllers.GetUsersController)
	e.GET("/masjids/:id", controllers.GetUserController)
	e.POST("/create-masjid", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)
	e.GET("/inventoris", controllers.GetInventorisController)
	e.GET("/inventoris/:id", controllers.GetInventoriController)
	e.GET("/inventoris", controllers.GetInventorisController)
	e.GET("/inventoris/:id", controllers.GetInventoriController)
	e.GET("/transaksis", controllers.GetTransaksisController)
	e.GET("/transaksis/:id", controllers.GetTransaksiController)

	eJWT := e.Group("")
	eJWT.Use(mid.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eJWT.DELETE("/masjids/:id", controllers.DeleteUserController)
	eJWT.PUT("/masjids/:id", controllers.UpdateUserController)
	eJWT.POST("/inventori-create/:id", controllers.CreateInventoriController)
	eJWT.DELETE("/inventoris/:id", controllers.DeleteInventoriController)
	eJWT.PUT("/inventoris", controllers.UpdateInventoriController)
	eJWT.POST("/transaksi-create/:id", controllers.CreateTransaksiController)
	eJWT.DELETE("/transaksi/:id", controllers.DeleteTransaksiController)
	eJWT.PUT("/transaksi", controllers.UpdateTransaksiController)
	return e
}
