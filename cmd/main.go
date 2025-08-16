package main

import (
	database "eventHub.com/internal/infra/DataBase"
	"eventHub.com/internal/infra/DataBase/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//inicializa o banco
	db, err := database.Connection()
	if err != nil{
		panic("Error with Database connection")
	}

	//inicializa o gin
	r := gin.Default()

	routes.SetUpRoutes(r, db)

	if err := r.Run(":8080"); err != nil{
		panic("error initializing database")
	}
}