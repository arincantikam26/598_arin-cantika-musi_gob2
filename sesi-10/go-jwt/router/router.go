package router

import (
	"go-jwt/controllers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middleware.Authenctication())
		productRouter.POST("/", controllers.CreateProduct())

		productRouter.PUT(":/productId", middleware.ProductAuthorization(), controllers.UpdateProduct)
	}
	return r
}
