package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	  "todo-API-with-go/models"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	dbPort   = "3306"
	dbname   = "todos_db"
	DB       *gorm.DB
	err      error
)

func InitDB() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, password, host, dbPort, dbname)
	dsn := config
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("====successfully connected to database=====")

}

func InitialMigration() {
	DB.AutoMigrate(&models.Todo{})
}

func GetDB() *gorm.DB {
	return DB
}