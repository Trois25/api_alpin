package models

import "gorm.io/gorm"

type Penerbit struct {
	gorm.Model

	Nama string `json:"nama" form:"nama"`

	Alamat string `json:"alamat" form:"alamat"`

	Kota string `json:"kota" form:"kota"`

	Telepon string `json:"telepon" form:"telepon"`
	
}
