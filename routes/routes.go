package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marialuizaleitao/gin-rest-api/controllers"
	"github.com/marialuizaleitao/gin-rest-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HandleRequest sets up the Gin router and defines routes
func HandleRequest() {
	r := gin.Default()

	// Hello message route
	r.GET("/:name", controllers.DisplayHelloMessage)

	// Members routes
	r.GET("/members", controllers.DisplayAllMembers)
	r.POST("/members", controllers.InsertMember)
	r.GET("/members/:id", controllers.FindMemberByID)
	r.DELETE("/members/:id", controllers.DeleteMember)
	r.PATCH("/members/:id", controllers.UpdateMember)

	// Swagger route
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
