package routers

import (
	"MyGram/controllers"
	"MyGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {

	router := gin.Default()

	authR := router.Group("/users")
	{
		authR.POST("/register", controllers.RegisterUser)
		authR.POST("/login", controllers.Login)
	}

	userR := router.Group("/users")
	{
		userR.Use(middlewares.JWTAuth())
		userR.PUT("/", controllers.UpdateUser)
		userR.DELETE("/", controllers.DeleteUser)
	}

	photoR := router.Group("/photos")
	{
		photoR.Use(middlewares.JWTAuth())
		photoR.POST("/", controllers.PostPhoto)
		photoR.GET("/", controllers.GetPhotos)
		photoR.PUT("/:id", controllers.UpdatePhoto)
		photoR.DELETE("/:id", controllers.DeletePhoto)
	}

	commentR := router.Group("/comments")
	{
		commentR.Use(middlewares.JWTAuth())
		commentR.POST("/", controllers.PostComment)
		commentR.GET("/", controllers.GetComments)
		commentR.PUT("/:id", controllers.UpdateComment)
		commentR.DELETE("/:id", controllers.DeleteComment)

	}
	sosmedR := router.Group("/socialmedias")
	{
		sosmedR.Use(middlewares.JWTAuth())
		sosmedR.POST("/", controllers.CreateSosmed)
		sosmedR.GET("/", controllers.GetSosmeds)
		sosmedR.PUT("/:id", controllers.UpdateSosmed)
		sosmedR.DELETE("/:id", controllers.DeleteSosmed)

	}

	return router
}
