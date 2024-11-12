package main

import (
	"github.com/gin-gonic/gin"

	"tierlist/database"
	"tierlist/routes"
)

func main() {
	database.ConnectDatabase()
	database.SeedUsers()
	database.SeedTierlists()
	r := gin.Default()
	routes.SetupUserRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }