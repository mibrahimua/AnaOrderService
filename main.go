package main

import (
	"AnaOrderService/config"
	"AnaOrderService/controller"
	"AnaOrderService/docs"
	"AnaOrderService/repository"
	"AnaOrderService/service"
	_ "fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func main() {
	db := config.GetDB()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewOrderService(productRepository)
	productController := controller.NewOrderController(productService)

	docs.SwaggerInfo.Title = "Ana Store - Environment: "
	docs.SwaggerInfo.Description = "API for product service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("swagger_host")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define your routes
	router.POST("/product", productController.CheckoutItems)

	// Start the server
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
