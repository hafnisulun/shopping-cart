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

	cart := models.Cart{}
	json.NewDecoder(w.Body).Decode(&cart)

	a.Equal(http.StatusCreated, w.Code)
}
