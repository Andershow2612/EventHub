package database

import (
	"log"
	"os"

	"eventHub.com/internal/entity"
	utils "eventHub.com/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	
	connection := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Connection error: " + err.Error())
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Address{},
		&entity.Category{},
		&entity.Event{},
		&entity.Kind{},
		&entity.Mode{},
		&entity.Role{},
		&entity.Status{},
		&entity.TicketKind{},
		&entity.Ticket{},
	)
	
	//para inputar os dados primarios de users
	utils.SeedRoles(db)

	if err != nil{
		log.Fatal("Migration error", err)
	}

	return db, nil
}