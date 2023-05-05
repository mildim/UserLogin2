package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mildim/UserLogin2/middleware"
	"github.com/mildim/UserLogin2/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
