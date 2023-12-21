package models

type Book struct {
	ID string `json:"id" form:"id" gorm:"primary_key"`

	Kategori string `json:"kategori" form:"kategori"`

	NamaBuku string `json:"nama_buku" form:"nama_buku"`

	Harga int `json:"harga" form:"harga"`

	Stock int `json:"stock" form:"stock"`

	PenerbitID string     `json:"penerbit_id" form:"penerbit_id"`
	Penerbit   Penerbit `json:"penerbit" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
