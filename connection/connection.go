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
	dsn := "root:DmAB1G1aSLMrGTkdtG4M@tcp(containers-us-west-41.railway.app:6955)/railway?parseTime=true"
	// dsn := "root:45218450@tcp(127.0.0.1:3306)/quotes?parseTime=true"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to open database connection")
	}
	// log.Println(DB.Exec("SELECT VERSION()"))
	log.Println("Database connection established")

	log.Println("Migrating database...")
	DB.AutoMigrate(models.Quote{})
	DB.AutoMigrate(models.User{})
	log.Println("Migrated successfully ")
}
