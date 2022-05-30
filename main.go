package main

import (
	"assignment2/config"
	"assignment2/controllers"
	"assignment2/repositories"
	"assignment2/services"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDB()

	router := gin.Default()

	orderRepo := repositories.NewOrderRepo(db)
	itemRepo := repositories.NewItemRepo(db)
	orderService := services.NewOrderService(orderRepo)
	itemService := services.NewItemService(itemRepo)
	orderController := controllers.NewOrderController(orderService, itemService)

	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders", orderController.GetOrders)
	router.DELETE("/orders/:orderId", orderController.DeleteOrders)
	router.PUT("/orders/:orderId", orderController.UpdateOrders)

	router.Run(config.APP_PORT)
}
