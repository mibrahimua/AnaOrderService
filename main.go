package main

import (
	"AnaOrderService/config"
	"AnaOrderService/controller"
	"AnaOrderService/docs"
	"AnaOrderService/repository"
	"AnaOrderService/service"
	"AnaOrderService/task"
	_ "fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func main() {
	db := config.GetDB()

	orderRepository := repository.NewOrderRepository(db)
	productService := service.NewOrderService(orderRepository)
	productController := controller.NewOrderController(productService)

	docs.SwaggerInfo.Title = "Ana Store - Environment: "
	docs.SwaggerInfo.Description = "API for order service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("swagger_host")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()

	go task.SyncReleasedUnusedStocks(orderRepository)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define your routes
	router.POST("/order/checkout", productController.CheckoutItems)

	// Start the server
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
