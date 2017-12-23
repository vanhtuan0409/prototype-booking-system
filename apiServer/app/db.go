package app

import (
	"apiServer/utils"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initializeDb() (*gorm.DB, error) {
	user := utils.GetEnv("DB_USER", "booking")
	password := utils.GetEnv("DB_PASSWORD", "password")
	host := utils.GetEnv("DB_HOST", "localhost:5432")
	database := utils.GetEnv("DB_DATABASE", "booking")
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		user,
		password,
		host,
		database,
	)
	fmt.Println(connStr)
	return gorm.Open("postgres", connStr)
}
