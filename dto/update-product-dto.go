package dto

type UpdateProductDTO struct {
	Harga     int    `json:"harga" form:"harga" binding:"required"`
	Stok      int    `json:"stok" form:"stok" binding:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" binding:"required"`
}
