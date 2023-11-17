package database

import (
	"final_project_1/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "12345"
	dbname   = "final_project_1"
)

var db *gorm.DB

func HandleDatabaseConnection() {
	psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = gorm.Open(postgres.Open(psqlinfo), &gorm.Config{})

	if err != nil {
		panic("failed connect to database")
	}

	db.AutoMigrate(entity.Todo{})

}

func GetDatabaseInstance() *gorm.DB {
	return db
}
