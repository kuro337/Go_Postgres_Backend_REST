package database 

import (
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	"fmt"
	"main/models"
)

var DB *gorm.DB 

func Connect() {
	dsn := "host=localhost user=testuser password=password dbname=firstdb port=5432 sslmode=disable TimeZone=America/New_York"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!= nil {
		panic("no db connection")
	}

	// pointer to DB
	DB = db 
	
	fmt.Println(db)

	// Adds the tables to our DB
	db.AutoMigrate(&models.User{} , &models.Role{} , &models.Permission{} , &models.Product{} , &models.Order{} , &models.OrderItem{})
}