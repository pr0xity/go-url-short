package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Shrt struct {
	ID           int    `json:"id"`
	Original_url string `json:"original_url" gorm:"not null"`
	Shrt         string `json:"shrt" gorm:"unique;not null"`
	Clicked      int    `json:"clicked" gorm:"not null"`
	Random       bool   `json:"random_url" gorm:"not null"`
	// CreatedAt    time.Time `json:"created_at"`
	// UpdatedAt    time.Time `json:"updated_at"`
	// DeletedAt    time.Time `json:"deleted_at"`
}

func Setup(host, port, username, password, database, timezone string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		host, port, username, password, database, timezone)
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database with error: " + err.Error())
	}
	err = db.AutoMigrate(&Shrt{})
	if err != nil {
		fmt.Println(err)
	}
}
