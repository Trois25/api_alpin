package controllers

import (
	"praktikum/config"
	"praktikum/helpers"
	"praktikum/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all books

func GetPenerbitController(c echo.Context) error {

	var penerbit []models.Penerbit

	if err := config.DB.Find(&penerbit).Error; err != nil {

		return c.JSON(http.StatusBadRequest,helpers.FailedResponse("error get data"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get all penerbit",penerbit))

}

// get book by id

func GetPenerbitSpesifikController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	penerbit := models.Penerbit{}
	err := config.DB.First(&penerbit, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest,helpers.FailedResponse("penerbit not found"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get penerbit",penerbit))
}

// create new book

func CreatePenerbitController(c echo.Context) error {

	penerbit := models.Penerbit{}

	c.Bind(&penerbit)

	if err := config.DB.Save(&penerbit).Error; err != nil {

		return c.JSON(http.StatusBadRequest,helpers.FailedResponse("Failed create penerbit"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success create new penerbit",penerbit))

}

// delete book by id

func DeletePenerbitController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	penerbit := models.Penerbit{}
	err := config.DB.Delete(&penerbit, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest,helpers.FailedResponse("Failed delete penerbit"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success delete penerbit",penerbit))

}

// update book by id

func UpdatePenerbitController(c echo.Context) error {

	// your solution here
	id,_ := strconv.Atoi(c.Param("id"))

	penerbit := models.Penerbit{}
	err := config.DB.First(&penerbit,id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest,helpers.FailedResponse("Failed update penerbit"))
	}

	update := new(models.Penerbit)
	if err := c.Bind(update); err != nil {
		return c.JSON(http.StatusBadRequest,helpers.FailedResponse("penerbit data is not valid"))
	}

	penerbit.Nama = update.Nama
	penerbit.Alamat = update.Alamat
	penerbit.Kota = update.Kota
	penerbit.Telepon = update.Telepon

	config.DB.Save(&penerbit)

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success update penerbit",penerbit))

}