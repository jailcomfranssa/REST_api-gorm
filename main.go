package main

import (
	"fmt"
	"github.com/jailcomfranssa/REST_api-gorm/database"
	"github.com/jailcomfranssa/REST_api-gorm/routes"
)

func main() {
	
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
