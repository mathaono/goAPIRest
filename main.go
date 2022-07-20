package main

import (
	"fmt"
	"goAPIRest/database"
	"goAPIRest/models"
	"goAPIRest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaBanco()
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
