package config

import(
	"praktikum/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := models.Config{

		DB_Username: "root",

		DB_Password: "",

		DB_Port: "3306",

		DB_Host: "localhost",

		DB_Name: "crud_go",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		config.DB_Username,

		config.DB_Password,

		config.DB_Host,

		config.DB_Port,

		config.DB_Name,
	)

	var e error
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	InitialMigration()

}

func InitialMigration() {

	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Penerbit{})

}