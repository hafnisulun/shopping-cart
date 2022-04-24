package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/shopping-cart/models"
)

type CartController struct{}

// POST /carts
// Create a cart
func (r CartController) Create(c *gin.Context) {
	cart := models.Cart{}
	result := models.DB.Create(&cart)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Create cart failed"})
	}
	c.JSON(http.StatusCreated, models.CartResponse{Data: cart})
}

// GET /carts/:cart_uuid
// Find a cart by UUID
func (r CartController) FindOne(c *gin.Context) {
	var cart models.Cart

	// Find the cart
	if err := models.DB.Preload("Items.Product").
		Where("uuid = ?", c.Param("cart_uuid")).
		First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, models.CartResponse{Data: cart})
}

// POST /carts/:cart_uuid/items
// Create a cart item
func (r CartController) CreateItem(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, http.StatusNotImplemented)
}
