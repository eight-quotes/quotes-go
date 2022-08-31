package connection

import (
	"log"

	"github.com/ochoad24/quotes-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	dsn := "root:j1pzxt3s0hDhz4h1MrRq@tcp(containers-us-west-60.railway.app:6561)/railway"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to open database connection")
	}
	log.Println("Database connection established")

	log.Println("Migrating database...")
	DB.AutoMigrate(models.Quote{})
	DB.AutoMigrate(models.User{})
	log.Println("Migrated successfully ")
}
