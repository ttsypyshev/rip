package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"rip/src/backend"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(backend.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&backend.Lang{}, &backend.Project{}, &backend.File{})
	if err != nil {
		panic("cant migrate db")
	}
}
