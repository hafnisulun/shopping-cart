package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartController struct{}

// POST /carts
// Create a cart
func (r CartController) Create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, http.StatusNotImplemented)
}

// GET /carts/:cart_uuid
// Find a cart by UUID
func (r CartController) FindOne(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, http.StatusNotImplemented)
}

// POST /carts/:cart_uuid/items
// Create a cart item
func (r CartController) CreateItem(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, http.StatusNotImplemented)
}
