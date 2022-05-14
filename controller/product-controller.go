package controller

import (
	"Ecommerce-Product/dto"
	"Ecommerce-Product/helper"
	"Ecommerce-Product/model"
	"Ecommerce-Product/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController interface {
	GetProducts(context *gin.Context)
	GetProduct(context *gin.Context)
	GetProductsByIDSeller(context *gin.Context)
	CreateProduct(context *gin.Context)
	UpdateProduct(context *gin.Context)
	DeleteProduct(context *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (p *productController) GetProducts(context *gin.Context) {
	products := p.productService.GetProducts()
	// res := helper.BuildResponse(true, "Products retrieved successfully", products)
	context.JSON(http.StatusOK, products)
}

func (p *productController) GetProduct(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid product id", "Error", model.Product{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	product := p.productService.GetProduct(idInt)
	// res := helper.BuildResponse(true, "Product retrieved successfully", product)
	context.JSON(http.StatusOK, product)
}

func (p *productController) GetProductsByIDSeller(context *gin.Context) {
	idSeller := context.Param("id")
	idSellerInt, err := strconv.Atoi(idSeller)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid seller id", "Error", model.Product{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	products := p.productService.GetProductsByIDSeller(idSellerInt)
	// res := helper.BuildResponse(true, "Products retrieved successfully", products)
	context.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(context *gin.Context) {
	var newProductDTO dto.NewProductDTO
	err := context.ShouldBind(&newProductDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid product data", err.Error(), model.Product{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	newProduct := p.productService.CreateProduct(newProductDTO)
	if newProduct.ID == 0 {
		res := helper.BuildErrorResponse("Error creating product", "Error", model.Product{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "Product created successfully", newProduct)
	context.JSON(http.StatusOK, res)
}

func (p *productController) UpdateProduct(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)

	var updateProductDTO dto.UpdateProductDTO
	err = context.ShouldBind(&updateProductDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid product data", "Error", model.Product{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	updatedProduct := p.productService.UpdateProduct(idInt, updateProductDTO)
	if updatedProduct.ID == 0 {
		res := helper.BuildErrorResponse("Error updating product", "Error", model.Product{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "Product updated successfully", updatedProduct)
	context.JSON(http.StatusOK, res)
}

func (p *productController) DeleteProduct(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		res := helper.BuildErrorResponse("Invalid product id", "Error", model.Product{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	productDeleted := p.productService.DeleteProduct(idInt)
	res := helper.BuildResponse(true, "Product deleted successfully", productDeleted)
	context.JSON(http.StatusOK, res)
}
