package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Kategori string `json:"kategori" form:"kategori"`

	NamaBuku string `json:"nama_buku" form:"nama_buku"`

	Harga int `json:"harga" form:"harga"`

	Stock int `json:"stock" form:"stock"`

	PenerbitID uint     `json:"penerbit_id" form:"penerbit_id"`
	Penerbit   Penerbit `json:"penerbit" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
