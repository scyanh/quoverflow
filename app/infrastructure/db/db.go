package db

import (
	"fmt"
	"go.uber.org/zap"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func Init() {

	/// PATH RELATIVE .env
	godotenv.Load(os.ExpandEnv(".env"))

	var err error
	DATABASE_USER := os.Getenv("DATABASE_USER")
	DATABASE_PASS := os.Getenv("DATABASE_PASS")
	DATABASE_URL := os.Getenv("DATABASE_URL")
	DATABASE_PORT := os.Getenv("DATABASE_PORT")
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		DATABASE_USER,
		DATABASE_PASS,
		DATABASE_URL,
		DATABASE_PORT,
		DATABASE_NAME)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		zap.S().Error("Failed to connect to db")
		panic(err)
	}
	zap.S().Infow("Database connected")

}

func GetDb() *gorm.DB {
	return db
}

func CloseDb() {
	db.Close()
}
