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
	r.DELETE("/members/:id", controllers.DeleteMember)
	r.PATCH("/members/:id", controllers.UpdateMember)
	r.Run()
}
