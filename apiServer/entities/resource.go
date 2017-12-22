package entities

import (
	"github.com/jinzhu/gorm"
)

type Resource struct {
	gorm.Model
	Availability bool
}
