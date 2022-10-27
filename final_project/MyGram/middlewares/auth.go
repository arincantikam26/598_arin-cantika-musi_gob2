package middlewares

import (
	"MyGram/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		log.Println(verifyToken)

		if err != nil {
			errors := helpers.FormatValidationError(err)
			errMessage := gin.H{"errors": errors}

			response := helpers.APIResponse("Token invalid", http.StatusBadRequest, "error", errMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		c.Set("currentUser", verifyToken)
		c.Next()
	}
}
