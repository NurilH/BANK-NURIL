package routes

import (
	"test-b/controllers"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// route users tanpa JWT
	e.POST("/nasabah", controllers.CreateNasabahControllers)
	e.GET("/nasabah", controllers.GetAllNasabahControllers)
	e.GET("/nasabah/:param", controllers.GetByKtpNasabahControllers)
	e.DELETE("/nasabah/:param", controllers.DeleteByIdNasabahControllers)
	e.PUT("/nasabah/:param", controllers.UpdateNasabahControllers)

	return e
}
