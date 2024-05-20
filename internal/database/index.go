package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(os.Getenv("string_mysql")), &gorm.Config{})
}

var DB *gorm.DB

func Initialize() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	dbConn, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	DB = dbConn

	return DB
}

func Close() error {
	sqlDB, err := DB.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}
