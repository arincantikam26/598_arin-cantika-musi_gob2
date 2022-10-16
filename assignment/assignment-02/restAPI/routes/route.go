package routes

import (
	config "restAPI/config"
	"restAPI/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func StartServer() *gin.Engine {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	routes := gin.Default()

	routes.POST("/orders", inDB.CreateOrders)
	routes.GET("/orders", inDB.GetOrders)

	routes.PUT("/orders/:orderId", inDB.UpdateOrder)

	// DeleteOrder
	routes.DELETE("/orders/:orderId", inDB.DeleteOrder)

	return routes
}
