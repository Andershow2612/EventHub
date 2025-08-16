package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() (*gorm.DB, error) {
	//conexão com o banco.
	connection := "host=localhost user=postgres password=root dbname=EventHub port=5432 sslmode=disable"

	//abre a conexão
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//valida se deu erro
	if err != nil {
		panic("Connection error")
	}

	return db, nil
}