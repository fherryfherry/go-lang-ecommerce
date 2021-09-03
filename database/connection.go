package database

import (
	"ecommerce/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start() (*gorm.DB, error) {
	// Connect the database
	hostname := "127.0.0.1"
	port	 := "3306"
	username := "root"
	password := ""
	database := "ecommerce"


	dsn := username + ":" + password + "@tcp("+hostname+":"+port+")/"+database+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func AutoMigrate() {
	db, err := Start()
	if err != nil {
		panic(err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&models.Users{}, &models.Tokens{}, &models.Categories{}, &models.Products{})
}

func Connect(c *gin.Context) *gorm.DB {

	db, err := Start()

	if err != nil {
		c.JSON(500, gin.H{"status":0,"message":"Can't connect to the database!"})
		c.Abort()
		return nil
	}

	return db
}

