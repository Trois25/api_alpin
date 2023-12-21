package models

type Penerbit struct {
	ID string `json:"id" form:"id" gorm:"primary_key"`

	Nama string `json:"nama" form:"nama"`

	Alamat string `json:"alamat" form:"alamat"`

	Kota string `json:"kota" form:"kota"`

	Telepon string `json:"telepon" form:"telepon"`
}
