package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.ConnectDatabase()

	fmt.Println("Iniciando o servidor Rest com o Go")

	routes.HandleRequest()
}
