package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tierlist/database"
	"tierlist/models"
)

func SetupTierlistRoutes (r *gin.Engine) {
	tierlist := r.Group("/tierlist")

	tierlist.GET("/", getAllTierlists)
	tierlist.GET("/:id", getTierlistById)
	tierlist.POST("/")
	tierlist.PUT("/:id")
	tierlist.DELETE("/:id")
}

func getAllTierlists(c *gin.Context) {
	var tierlists []models.Tierlist
	if err := database.DB.Find(&tierlists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
	}
	c.JSON(http.StatusOK, tierlists)
}

func getTierlistById(c *gin.Context) {
	var tierlists []models.Tierlist
	id := c.Param("id")
	if err := database.DB.First(&tierlists, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Tierlist not found"})
	}
	c.JSON(http.StatusOK, tierlists)
}