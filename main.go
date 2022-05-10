package main

import (
	"Ecommerce-Product/conn"
	"Ecommerce-Product/controller"
	"Ecommerce-Product/repository"
	"Ecommerce-Product/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = conn.SetupDatabaseConnection()
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	productService    service.ProductService       = service.NewProductService(productRepository)
	productController controller.ProductController = controller.NewProductController(productService)
)

func main() {
	defer conn.CloseDatabaseConnection(db)
	router := gin.Default()

	routes := router.Group("/api/product")
	{
		routes.GET("/", productController.GetProducts)
		routes.GET("/:id", productController.GetProduct)
		routes.GET("/seller/:id", productController.GetProductsByIDSeller)
		routes.POST("/", productController.CreateProduct)
		routes.PUT("/:id", productController.UpdateProduct)
		routes.DELETE("/:id", productController.DeleteProduct)
	}

	err := router.Run("192.168.100.8:8081")
	if err != nil {
		return
	}
}
