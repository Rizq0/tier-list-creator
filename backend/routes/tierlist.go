package routes

import "github.com/gin-gonic/gin"

func SetupTierlistRoutes (r *gin.Engine) {
	tierlist := r.Group("/tierlist")

	tierlist.GET("/")
	tierlist.GET("/:id")
	tierlist.POST("/")
	tierlist.PUT("/:id")
	tierlist.DELETE("/:id")
}