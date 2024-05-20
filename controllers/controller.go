package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DisplayHelloMessage(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name + "! Are you a Beatles fan?",
	})
}

func DisplayAllMembers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ID": 1, "Name": "George Harrison", "Role": "Guitar, Vocals"})
}
