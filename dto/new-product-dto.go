package dto

type NewProductDTO struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Harga     int    `json:"harga" form:"harga" binding:"required"`
	Stok      int    `json:"stok" form:"stok" binding:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" binding:"required"`
	Gambar    string `json:"gambar" form:"gambar" binding:"required"`
	IDPenjual uint64 `json:"id_penjual" form:"id_penjual" binding:"required"`
}
