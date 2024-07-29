package db

import (
	"fmt"

	"github.com/macreai/todo-app-clean-architecture/configs"
	"github.com/macreai/todo-app-clean-architecture/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	config := configs.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.ActivityUser{})
	return db, nil
}
