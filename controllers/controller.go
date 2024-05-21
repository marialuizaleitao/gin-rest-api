package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marialuizaleitao/gin-rest-api/database"
	"github.com/marialuizaleitao/gin-rest-api/models"
)

// DisplayHelloMessage godoc
// @Summary Displays a hello message
// @Description Displays a hello message with the provided name
// @Produce json
// @Param name path string true "Name of the person"
// @Success 200 {string} string "message: Hello {name}! Are you a Beatles fan?"
// @Router /{name} [get]
func DisplayHelloMessage(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name + "! Are you a Beatles fan?",
	})
}

// DisplayAllMembers godoc
// @Summary Displays all members
// @Description Retrieves a list of all members
// @Produce json
// @Success 200 {array} models.Member "List of members"
// @Router /members [get]
func DisplayAllMembers(c *gin.Context) {
	var members []models.Member
	database.DB.Find(&members)
	c.JSON(http.StatusOK, members)
}

// InsertMember godoc
// @Summary Inserts a new member
// @Description Inserts a new member with the provided data
// @Produce json
// @Accept json
// @Param member body models.Member true "Member data"
// @Success 200 {object} models.Member "Inserted member data"
// @Router /members [post]
func InsertMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateMemberData(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&member)
	c.JSON(http.StatusOK, member)
}

// FindMemberByID godoc
// @Summary Retrieves a member by ID
// @Description Retrieves a member with the specified ID
// @Produce json
// @Param id path string true "Member ID"
// @Success 200 {object} models.Member "Member data"
// @Router /members/{id} [get]
func FindMemberByID(c *gin.Context) {
	var member models.Member
	id := c.Params.ByName("id")

	if err := database.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}
	c.JSON(http.StatusOK, member)
}

// DeleteMember godoc
// @Summary Deletes a member by ID
// @Description Deletes a member with the specified ID
// @Produce json
// @Param id path string true "Member ID"
// @Success 200 {string} string "message: Member deleted"
// @Router /members/{id} [delete]
func DeleteMember(c *gin.Context) {
	var member models.Member
	id := c.Params.ByName("id")

	database.DB.Delete(&member, id)
	c.JSON(http.StatusOK, gin.H{"message": "Member deleted"})
}

// UpdateMember godoc
// @Summary Updates a member by ID
// @Description Updates a member with the specified ID
// @Produce json
// @Accept json
// @Param id path string true "Member ID"
// @Param member body models.Member true "Member data"
// @Success 200 {object} models.Member "Updated member data"
// @Router /members/{id} [patch]
func UpdateMember(c *gin.Context) {
	var member models.Member
	id := c.Params.ByName("id")
	database.DB.First(&member, id)
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateMemberData(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&member).UpdateColumns(member)
	c.JSON(http.StatusOK, member)
}
