package controllers

import (
	"math/rand"
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
	err := config.DB.Preload("Penerbit").First(&book, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Book not found"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get book", book))
}

// create new book

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	// Validate PenerbitID
	if book.PenerbitID == 0 {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Penerbit ID is required"))
	}

	// Check if the specified Penerbit exists
	var penerbit models.Penerbit
	err := config.DB.First(&penerbit, book.PenerbitID).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Invalid Penerbit ID"))
	}

	// Generate a custom string ID
	customID := generateCustomID(book.Kategori)
	book.ID = customID

	// Save the book
	if err := config.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to create book"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success create new book"))
}

// Function to generate custom string ID
func generateCustomID(kategori string) string {
	if len(kategori) == 0 {
        // Handle the case where kategori is empty
        return ""  // or some default value
    }
	// Assume kategori is not empty
	firstLetter := string(kategori[0])

	// Generate a random integer (you can replace this with your own logic)
	randomInt := rand.Intn(10000)

	// Concatenate the first letter and random integer
	return firstLetter + strconv.Itoa(randomInt)
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
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success delete book", book))

}

// update book by id

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	err := config.DB.First(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed update book"))
	}

	update := new(models.Book)
	if err := c.Bind(update); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Book data is not valid"))
	}

	// Validate PenerbitID
	if update.PenerbitID != 0 {
		// Check if the specified Penerbit exists
		var penerbit models.Penerbit
		err := config.DB.First(&penerbit, update.PenerbitID).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Invalid Penerbit ID"))
		}
		// Update Penerbit information
		book.PenerbitID = update.PenerbitID
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
