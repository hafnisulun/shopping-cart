package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"
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

	// Add "Alexa Speaker" to the cart (1 of 3)
	productUUID, err := uuid.FromString("6bdd17e6-3ed9-4ccd-a69a-1afe1883beca")
	if err != nil {
		a.Error(err)
	}

	item := models.CartItemInput{
		ProductUUID: productUUID,
	}

	reqBody, err := json.Marshal(item)
	if err != nil {
		a.Error(err)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v1/carts/"+cart.UUID.String()+"/items", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	a.Equal(http.StatusCreated, w.Code)

	// Add "Alexa Speaker" to the cart (2 of 3)
	productUUID, err = uuid.FromString("6bdd17e6-3ed9-4ccd-a69a-1afe1883beca")
	if err != nil {
		a.Error(err)
	}

	item = models.CartItemInput{
		ProductUUID: productUUID,
	}

	reqBody, err = json.Marshal(item)
	if err != nil {
		a.Error(err)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v1/carts/"+cart.UUID.String()+"/items", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	a.Equal(http.StatusCreated, w.Code)

	// Add "Alexa Speaker" to the cart (3 of 3)
	productUUID, err = uuid.FromString("6bdd17e6-3ed9-4ccd-a69a-1afe1883beca")
	if err != nil {
		a.Error(err)
	}

	item = models.CartItemInput{
		ProductUUID: productUUID,
	}

	reqBody, err = json.Marshal(item)
	if err != nil {
		a.Error(err)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v1/carts/"+cart.UUID.String()+"/items", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	a.Equal(http.StatusCreated, w.Code)

	// Get the cart
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/carts/"+cart.UUID.String(), nil)
	router.ServeHTTP(w, req)

	finalCartResponse := models.CartResponse{}
	json.NewDecoder(w.Body).Decode(&finalCartResponse)
	finalCart := finalCartResponse.Data

	a.Equal(http.StatusOK, w.Code)
	a.Equal(295.65, finalCart.Total)
}
