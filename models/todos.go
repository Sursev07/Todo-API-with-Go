package models
import (
	"gorm.io/gorm"
	// "time"
)

type Todo struct {
	gorm.Model
	ID int `gorm:"default:uuid_generate_v3()"`
	Title string `json: "title"`
	Description string `json: "description"`
	// DueDate time.Time `json:"due_date"`
	Category string `json:"category"`
}

// Result is an array of product
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}



