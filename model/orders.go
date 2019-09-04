package model

import (
	"github.com/jinzhu/gorm"
	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type Orders struct {
	gorm.Model
	User     User
	Products []OrdersList
}

type OrdersList struct {
	gorm.Model
	Count   int
	Product Product
}
