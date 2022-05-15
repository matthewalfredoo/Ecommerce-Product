package repository

import (
	"Ecommerce-Product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() []model.Product
	GetProduct(id int) model.Product
	GetProductsByIDSeller(idSeller int) []model.Product
	CreateProduct(product model.Product) model.Product
	UpdateProduct(id int, product model.Product) model.Product
	DeleteProduct(id int) model.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(conn *gorm.DB) ProductRepository {
	return &productConnection{
		connection: conn,
	}
}

func (db *productConnection) GetProducts() []model.Product {
	var products []model.Product
	db.connection.Find(&products)
	return products
}

func (db *productConnection) GetProduct(id int) model.Product {
	var product model.Product
	db.connection.First(&product, id)
	return product
}

func (db *productConnection) GetProductsByIDSeller(idSeller int) []model.Product {
	var products []model.Product
	db.connection.Where("id_penjual = ?", idSeller).Find(&products)
	return products
}

func (db *productConnection) CreateProduct(product model.Product) model.Product {
	db.connection.Create(&product)
	return product
}

func (db *productConnection) UpdateProduct(id int, product model.Product) model.Product {
	var productUpdate model.Product
	db.connection.First(&productUpdate, id)

	if productUpdate.Harga != product.Harga {
		productUpdate.Harga = product.Harga
	}
	if productUpdate.Stok != product.Stok {
		productUpdate.Stok = product.Stok
	}
	if productUpdate.Deskripsi != product.Deskripsi {
		productUpdate.Deskripsi = product.Deskripsi
	}

	db.connection.Save(&productUpdate)
	return productUpdate
}

func (db *productConnection) DeleteProduct(id int) model.Product {
	var product model.Product
	db.connection.First(&product, id).Delete(&product)
	return product
}
