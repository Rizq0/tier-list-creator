package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	users := r.Group("/users")

	users.GET("/me", handleMe) 

}

func handleMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "joe",
	})
}