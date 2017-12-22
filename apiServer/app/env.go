package app

import (
	"apiServer/entities"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Env struct {
	DB *gorm.DB
}

func NewEnv() *Env {
	db, err := initializeDb()
	if err != nil {
		fmt.Println("ERR:", err)
		panic("Cannot connect to db")
	}
	migrateDb(db)

	return &Env{
		DB: db,
	}
}

func migrateDb(db *gorm.DB) {
	db.AutoMigrate(
		&entities.Resource{},
	)
}
