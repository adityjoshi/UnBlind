package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=aditya dbname=hospital port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Print("database connected successfully ⚡️")
	}

	// Migrate the schema
	DB.AutoMigrate(&Users{})
}

var DB *gorm.DB

type Users struct {
	User_id       uint   `gorm:"primaryKey"`
	ContactNumber string `json:"ContactNumber" gorm:"not null"`
	Email         string `json:"Email" gorm:"not null;unique"`
	Password      string `json:"Password"`
}
