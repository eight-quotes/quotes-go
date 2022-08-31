package connection

import (
	"log"

	"github.com/ochoad24/quotes-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(sqlite.Open("quotes.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to open database connection")
	}
	log.Println("Database connection established")

	log.Println("Migrating database...")
	DB.AutoMigrate(models.Quote{})
	DB.AutoMigrate(models.User{})
	log.Println("Migrated successfully ")
}
