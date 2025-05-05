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

func TestPanicRouter(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		fmt.Fprint(w, "Panic HttpRouter: ", err)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("Panic!")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic HttpRouter: Panic!", string(body))
}
