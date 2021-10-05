package models
import (
	// "gorm.io/gorm"
	"time"
)

type OwnModel struct {
    ID        uint       `gorm:"primary_key"`
    CreatedAt time.Time  `json:"-"`
    UpdatedAt time.Time  `json:"-"`
    DeletedAt *time.Time `json:"-";sql:"index"`
}

type Todo struct {
	OwnModel
	Title string `json: "title"`
	Description string `json: "description"`
	DueDate time.Time `json:"due_date"`
	Category string `json:"category"`
}

// Result is an array of product
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}


