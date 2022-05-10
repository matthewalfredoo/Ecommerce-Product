package model

import (
	"time"
)

type Product struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Nama      string    `gorm:"type:varchar(255)" json:"nama"`
	Harga     int       `gorm:"type:int" json:"harga"`
	Stok      int       `gorm:"type:int" json:"stok"`
	Deskripsi string    `gorm:"type:text" json:"deskripsi"`
	Gambar    string    `gorm:"type:varchar(255)" json:"gambar"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
	IDPenjual uint64    `gorm:"type:int" json:"id_penjual"`
}
