package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marialuizaleitao/gin-rest-api/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/members", controllers.DisplayAllMembers)
	r.GET("/:name", controllers.DisplayHelloMessage)
	r.POST("/members", controllers.InsertMember)
	r.GET("/members/:id", controllers.FindMemberByID)
	r.Run()
}
