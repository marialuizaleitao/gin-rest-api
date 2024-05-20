package main

import (
	"github.com/marialuizaleitao/gin-rest-api/database"
	"github.com/marialuizaleitao/gin-rest-api/models"
	"github.com/marialuizaleitao/gin-rest-api/routes"
)

func main() {
	database.Connect()
	models.Members = []models.Member{
		{ID: 1, Name: "John Lennon", Role: "Vocals, Guitar"},
		{ID: 2, Name: "Paul McCartney", Role: "Vocals, Bass"},
		{ID: 3, Name: "George Harrison", Role: "Guitar, Vocals"},
		{ID: 4, Name: "Ringo Starr", Role: "Drums, Vocals"},
	}
	routes.HandleRequest()
}
