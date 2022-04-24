package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hafnisulun/shopping-cart/models"
	"github.com/stretchr/testify/assert"
)

func TestCartWithMinQtyPromo(t *testing.T) {
	router := setupRouter()

	a := assert.New(t)

	// Create a cart
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/carts", nil)
	router.ServeHTTP(w, req)

	cartResponse := models.CartResponse{}
	json.NewDecoder(w.Body).Decode(&cartResponse)
	cart := cartResponse.Data

	a.Equal(http.StatusCreated, w.Code)

	// Get the cart
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/carts/"+cart.UUID.String(), nil)
	router.ServeHTTP(w, req)

	a.Equal(http.StatusOK, w.Code)
}
