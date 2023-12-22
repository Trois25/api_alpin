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

func GetBooksController(c echo.Context) error {

	var books []models.Book

	if err := config.DB.Preload("Penerbit").Find(&books).Error; err != nil {

		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error get data"))

	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get all books", books))

}

// get book by id

func GetBookController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	err := config.DB.First(&book, id).Preload("Penerbit").Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Book not found"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get book", book))
}

// create new book

func CreateBookController(c echo.Context) error {
    // Bind the JSON request body to the book structure
    book := models.Book{}
    if err := c.Bind(&book); err != nil {
        return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Invalid request payload"))
    }

    // Validate Penerbit name
    if book.Penerbit.Nama == "" {
        return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Nama Penerbit is required"))
    }

    // Check if the specified Penerbit exists based on the name
    var penerbit models.Penerbit
    err := config.DB.Where("nama = ?", book.Penerbit.Nama).First(&penerbit).Error
    if err != nil {
        return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Penerbit tidak ditemukan"))
    }

    // Associate the existing Penerbit with the book
    config.DB.Model(&book).Association("Penerbit").Append(&penerbit)

    // Save the book
    if err := config.DB.Save(&book).Error; err != nil {
        return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to create book"))
    }

    return c.JSON(http.StatusOK, helpers.SuccessResponse("Success create new book"))
}

// delete book by id

func DeleteBookController(c echo.Context) error {

	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	err := config.DB.Delete(&book, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed delete book"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success delete book"))

}

// update book by id

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	err := config.DB.Preload("Penerbit").First(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed update book"))
	}

	update := new(models.Book)
	if err := c.Bind(update); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Book data is not valid"))
	}

	// Validate Penerbit Name
	if update.Penerbit.Nama != "" {
		// Check if the specified Penerbit exists
		var penerbit models.Penerbit
		err := config.DB.Where("nama = ?", update.Penerbit.Nama).First(&penerbit).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Invalid Penerbit Name"))
		}
		// Update Penerbit information
		book.PenerbitID = penerbit.ID
		book.Penerbit = penerbit
	}

	// Update other book information
	book.Kategori = update.Kategori
	book.NamaBuku = update.NamaBuku
	book.Harga = update.Harga
	book.Stock = update.Stock

	config.DB.Save(&book)

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success update book"))
}
