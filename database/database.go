package database

import (
	"fmt"
	"gofiber/initializers"
	"gofiber/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnvVariables()
}

func ConnectToDB() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dbsetting := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dbsetting), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting DB")
	}
	log.Println("Connecting to DB success")

	// setiap program running, fungsi ini kan dipanggil terus... kyknya jadi migrate terus deh, mungkin lebih baik nnti untuk migrasi dipisah aja
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

}
