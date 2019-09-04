package model

import (
	"github.com/jinzhu/gorm"
	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Category    string `gorm:"type:varchar(50);not null"`
	Description string `gorm:"type:varchar(300);not null"`
	ImageUrl    string `gorm:"type:varchar(500);not null"`
	Count       int
	Price       int
}
