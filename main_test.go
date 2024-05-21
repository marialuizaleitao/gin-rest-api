package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marialuizaleitao/gin-rest-api/controllers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestHelloMessageStatusCode(t *testing.T) {
	r := setupTestRoutes()
	r.GET("/:name", controllers.DisplayHelloMessage)
	req, _ := http.NewRequest("GET", "/maria", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
