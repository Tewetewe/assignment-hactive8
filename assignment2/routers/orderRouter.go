package routers

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/order", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	return router
}
