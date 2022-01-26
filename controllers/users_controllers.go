package controllers

import (
	"net/http"
	"strconv"
	"test-b/databases"
	"test-b/models"

	response "test-b/responses"

	"github.com/labstack/echo/v4"
)

// controller untuk menambahkan Nasabah (registrasi) next
func CreateNasabahControllers(c echo.Context) error {
	new_nasabah := models.Nasabah{}
	c.Bind(&new_nasabah)

	data, err := databases.CreateNasabah(&new_nasabah)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func GetAllNasabahControllers(c echo.Context) error {
	data, e := databases.GetAllNasabah()
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func GetByKtpNasabahControllers(c echo.Context) error {
	ktp, er := strconv.Atoi(c.Param("param"))
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("false param"))
	}

	data, e := databases.GetByKtpNasabah(ktp)
	if e != nil || data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func DeleteByIdNasabahControllers(c echo.Context) error {
	Id, er := strconv.Atoi(c.Param("param"))
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("false param"))
	}

	data, e := databases.DeleteByIdNasabah(Id)
	if e != nil || data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Delete Success"))
}

func UpdateNasabahControllers(c echo.Context) error {
	new_nasabah := models.Nasabah{}
	c.Bind(&new_nasabah)
	Id, er := strconv.Atoi(c.Param("param"))
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("false param"))
	}

	data, err := databases.UpdateNasabah(Id, &new_nasabah)

	if err != nil || data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}
