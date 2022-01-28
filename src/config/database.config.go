package config

import (
	"fmt"
	"github.com/duchai27798/golang_api_tutorial/src/data/entity"
	"github.com/duchai27798/golang_api_tutorial/src/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		utils.LogObj("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.LogObj("Failed to create a connection to DB")
	} else {
		utils.LogObj(db, "connect successful")
	}
	// migrate db
	errBD := db.AutoMigrate(&entity.User{}, &entity.Book{})
	if errBD != nil {
		utils.LogObj(errBD, "error db")
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		utils.LogObj(err)
	}
	dbSQL.Close()
}
