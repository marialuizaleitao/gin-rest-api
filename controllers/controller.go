package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marialuizaleitao/gin-rest-api/database"
	"github.com/marialuizaleitao/gin-rest-api/models"
	"net/http"
)

func DisplayHelloMessage(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name + "! Are you a Beatles fan?",
	})
}

func DisplayAllMembers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Members)
}

func InsertMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&member)
	c.JSON(http.StatusOK, member)
}
