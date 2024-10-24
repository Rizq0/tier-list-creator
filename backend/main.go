package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tierlist/routes"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	routes.SetupUserRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }