package postgres

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("host=host.docker.internal user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return nil
	}

	db.AutoMigrate(&User{}, &Product{}, &Cart{}, &Transaction{}, &CartItem{})

	logrus.Info("PostgreSQL Connected Successfully")
	return db
}
