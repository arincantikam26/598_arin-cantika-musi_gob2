package routes

import (
	"live_code/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	router.POST("/arin", controllers.CreateMy)
	router.GET("/arin", controllers.GetAllMy)

	return router

}
