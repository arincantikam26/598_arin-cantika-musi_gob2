package routers

func StartServer() *gin.Engine {

	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	return router

}
