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

    // Generate a custom string ID
    customID := generateCustomID(book.Kategori)
    book.ID = customID

    // Set the PenerbitID based on the retrieved Penerbit's ID
    book.PenerbitID = penerbit.ID

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
		return "kategori tidak boleh kosong" // or some default value
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
