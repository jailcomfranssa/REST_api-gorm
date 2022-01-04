package database

import (
	"fmt"
	"github.com/jailcomfranssa/REST_api-gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaComBancoDeDados() *gorm.DB {
	stringDeConexao := "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados %s", err)
	}
	DB.AutoMigrate(&models.Personalidade{})

	fmt.Printf("conectado %s", DB)
	return DB
}
