package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParamPatternRouter(t *testing.T) {
	var bodyExpected string
	router := httprouter.New()
	router.GET("/products/:productId/items/:itemId", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		productId := params.ByName("productId")
		itemId := params.ByName("itemId")
		bodyExpected = "Product: " + productId + " Item: " + itemId
		fmt.Fprint(w, bodyExpected)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/3", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, bodyExpected, string(body))
}

func TestAllParamPatternRouter(t *testing.T) {
	var image string
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		image = params.ByName("image")
		fmt.Fprint(w, image)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/avatar.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, image, string(body))
}
