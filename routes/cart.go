package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/shopping-cart/controllers"
)

func Carts(router *gin.RouterGroup) {
	cartsGroup := router.Group("/carts")
	{
		cart := new(controllers.CartController)
		cartsGroup.POST("", cart.Create)
		cartsGroup.GET(":cart_uuid", cart.FindOne)
		cartsGroup.POST(":cart_uuid/items", cart.CreateItem)
	}
}
