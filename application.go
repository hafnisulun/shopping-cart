package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/shopping-cart/models"
	"github.com/hafnisulun/shopping-cart/routes"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	models.ConnectDatabase()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		routes.Carts(v1)
	}
	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
