package service

import (
	"Ecommerce-Product/dto"
	"Ecommerce-Product/model"
	"Ecommerce-Product/repository"
	"github.com/mashingan/smapping"
	"time"
)

type ProductService interface {
	GetProducts() []model.Product
	GetProduct(id int) model.Product
	GetProductsByIDSeller(idSeller int) []model.Product
	CreateProduct(dto dto.NewProductDTO) model.Product
	UpdateProduct(id int, dto dto.UpdateProductDTO) model.Product
	DeleteProduct(id int) model.Product
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: repository,
	}
}

func (ps *productService) GetProducts() []model.Product {
	return ps.productRepository.GetProducts()
}

func (ps *productService) GetProduct(id int) model.Product {
	return ps.productRepository.GetProduct(id)
}

func (ps *productService) GetProductsByIDSeller(idSeller int) []model.Product {
	return ps.productRepository.GetProductsByIDSeller(idSeller)
}

func (ps *productService) CreateProduct(dto dto.NewProductDTO) model.Product {
	productDTOToModel := model.Product{}

	err := smapping.FillStruct(&productDTOToModel, smapping.MapFields(&dto))
	if err != nil {
		return model.Product{}
	}

	productDTOToModel.Nama = dto.Nama
	productDTOToModel.Harga = dto.Harga
	productDTOToModel.Stok = dto.Stok
	productDTOToModel.Deskripsi = dto.Deskripsi
	productDTOToModel.Gambar = dto.Gambar
	productDTOToModel.IDPenjual = dto.IDPenjual

	productDTOToModel.CreatedAt = time.Now()
	productDTOToModel.UpdatedAt = time.Now()
	return ps.productRepository.CreateProduct(productDTOToModel)
}

func (ps *productService) UpdateProduct(id int, dto dto.UpdateProductDTO) model.Product {
	productDTOToModel := model.Product{}
	err := smapping.FillStruct(&productDTOToModel, smapping.MapFields(&dto))
	if err != nil {
		return model.Product{}
	}
	productDTOToModel.UpdatedAt = time.Now()
	return ps.productRepository.UpdateProduct(id, productDTOToModel)
}

func (ps *productService) DeleteProduct(id int) model.Product {
	return ps.productRepository.DeleteProduct(id)
}
