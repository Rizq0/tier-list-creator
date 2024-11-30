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
	tierlist.POST("/", createTierlist)
	tierlist.PUT("/:id", updateTierlist)
	tierlist.DELETE("/:id", deleteTierList)
}

func getAllTierlists(c *gin.Context) {
	var tierlists []models.Tierlist
	if err := database.DB.Find(&tierlists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
		return
	}
	c.JSON(http.StatusOK, tierlists)
}

func getTierlistById(c *gin.Context) {
	var tierlists []models.Tierlist
	id := c.Param("id")
	if err := database.DB.First(&tierlists, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Tierlist not found"})
		return
	}
	c.JSON(http.StatusOK, tierlists)
}

func createTierlist(c *gin.Context) {
	var request models.CreateTierlistRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	tierlist := models.Tierlist{
		Name: request.Name,
		Description: request.Description,
		CreatorID: request.Creator,
	}

	if err := database.DB.Create(&tierlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
		return
	}

	for _, tier := range request.Tiers {
		newTier := models.Tier{
			TierlistID: tierlist.ID,
			Text: tier.Name,
			Colour: tier.Colour,
		}
		if err := database.DB.Create(&newTier).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
			return
		}
	}

	for _, item := range request.Items {
		newItem := models.Item{
			TierlistID: tierlist.ID,
			Text: item.Text,
			Image: item.Image,
			TierText: item.Tier,
		}
		if err := database.DB.Create(&newItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
			return
		}
	}

	c.JSON(http.StatusCreated, tierlist)
}

func updateTierlist(c *gin.Context) {
	var request models.UpdateTierlistRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id := c.Param("id")
	var tierlist models.Tierlist
	if err := database.DB.First(&tierlist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Tierlist not found"})
		return
	}

	for _, tier := range request.Tiers {
		newTier := models.Tier{
			TierlistID: tierlist.ID,
			Text: tier.Name,
			Colour: tier.Colour,
		}
		if err := database.DB.Create(&newTier).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
			return
		}
	}

	for _, item := range request.Items {
		newItem := models.Item{
			TierlistID: tierlist.ID,
			Text: item.Text,
			Image: item.Image,
			TierText: item.Tier,
		}
		if err := database.DB.Create(&newItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
			return
		}
	}

	c.JSON(http.StatusCreated, tierlist)
}

func deleteTierList(c *gin.Context) {
	id := c.Param("id")
	var tierlist models.Tierlist
	if err := database.DB.First(&tierlist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Tierlist not found"})
		return
	}

	if err := database.DB.Delete(&tierlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
		return
	}

	if err := database.DB.Where("tierlist_id = ?", id).Delete(&models.Tier{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
		return
	}

	if err := database.DB.Where("tierlist_id = ?", id).Delete(&models.Item{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Tierlist deleted"})
}