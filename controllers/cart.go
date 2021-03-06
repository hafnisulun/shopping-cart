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
	// Create a cart
	cart := models.Cart{}
	if err := models.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Create cart failed"})
		return
	}

	// Send success response
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Calculate total
	cart.Total = 0

	for _, item := range cart.Items {
		var itemTotal float64 = item.Product.Price * float64(item.Qty)

		// Check promo "buy A get B"
		var buyAGetBPromo models.BuyAGetBPromo
		result := models.DB.
			Where("get_product_id = ?", item.ProductID).
			Where("get_qty <= ?", item.Qty).
			First(&buyAGetBPromo)

		if result.RowsAffected > 0 {
			if buyAGetBPromo.BuyProductID == buyAGetBPromo.GetProductID {
				if item.Qty >= buyAGetBPromo.BuyQty+buyAGetBPromo.GetQty {
					itemTotal = item.Product.Price * float64(item.Qty-buyAGetBPromo.GetQty)
				}
			} else {
				for _, scannedItem := range cart.Items {
					if scannedItem.ProductID == buyAGetBPromo.BuyProductID && scannedItem.Qty >= buyAGetBPromo.BuyQty {
						itemTotal = item.Product.Price * float64(item.Qty-buyAGetBPromo.GetQty)
					}
				}
			}
		}

		// Check promo "min qty"
		var minQtyPromo models.MinQtyPromo
		result = models.DB.
			Where("product_id = ?", item.Product.ID).
			Where("qty <= ?", item.Qty).
			First(&minQtyPromo)

		if result.RowsAffected > 0 {
			itemTotal -= itemTotal * float64(minQtyPromo.Discount) / 100
		}

		cart.Total += itemTotal
	}

	// Send success response
	c.JSON(http.StatusOK, models.CartResponse{Data: cart})
}

// POST /carts/:cart_uuid/items
// Create a cart item
func (r CartController) CreateItem(c *gin.Context) {
	// Validate input
	var input models.CartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body invalid"})
		return
	}

	// Get the cart
	var cart models.Cart
	if err := models.DB.
		Where("uuid = ?", c.Param("cart_uuid")).
		First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Get the product
	var product models.Product
	if err := models.DB.
		Where("uuid = ?", input.ProductUUID).
		First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Check if the product already in the cart
	var item models.CartItem
	result := models.DB.Preload("Product").
		Where("cart_id = ?", cart.ID).
		Where("product_id = ?", product.ID).
		Find(&item)

	if result.RowsAffected > 0 {
		// Add qty of the item
		item.Qty++
		if err := models.DB.Save(&item).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Add item qty failed"})
			return
		}
	} else {
		// Create a new item
		item = models.CartItem{
			CartID:    cart.ID,
			ProductID: product.ID,
			Product:   product,
			Qty:       1,
		}
		if err := models.DB.Create(&item).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Create item failed"})
			return
		}
	}

	// Send success response
	c.JSON(http.StatusCreated, models.CartItemResponse{Data: item})
}
