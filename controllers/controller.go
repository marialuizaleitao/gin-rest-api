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
	var members []models.Member
	database.DB.Find(&members)
	c.JSON(http.StatusOK, members)
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

func FindMemberByID(c *gin.Context) {
	var member models.Member
	id := c.Params.ByName("id")

	if err := database.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}
	c.JSON(http.StatusOK, member)
}

func DeleteMember(c *gin.Context) {
	var member models.Member
	id := c.Params.ByName("id")

	database.DB.Delete(&member, id)
	c.JSON(http.StatusOK, gin.H{"message": "Member deleted"})
}
