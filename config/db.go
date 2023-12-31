package config

import (
	"fmt"
	"synapsis-go-try/helpers"
	"synapsis-go-try/models"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Init() {
	e := godotenv.Load()
	if e != nil {
		helpers.Logger("error", "Error getting env")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDriver := os.Getenv("DB_DRIVER")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)

	conn, err := gorm.Open(dbDriver, dbURI)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
	}

	db = conn
	db.Debug().AutoMigrate(&models.User{}, &models.Item{}, &models.Payment{}, &models.Shoppingcart{})

	db.Model(&models.Shoppingcart{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Model(&models.Shoppingcart{}).AddForeignKey("item_id", "items(item_id)", "CASCADE", "CASCADE")
	db.Model(&models.Payment{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Model(&models.Payment{}).AddForeignKey("purchase_id", "shoppingcarts(purchase_id)", "CASCADE", "CASCADE")

}

func GetDB() *gorm.DB {
	return db
}
