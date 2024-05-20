package main

import (
	"github.com/marialuizaleitao/gin-rest-api/database"
	"github.com/marialuizaleitao/gin-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequest()
}
