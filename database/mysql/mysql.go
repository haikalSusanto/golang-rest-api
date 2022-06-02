package mysql

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("root:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_ROOT_PASSWORD"), os.Getenv("DB_NAME"))
	logrus.Info(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	logrus.Info("PostgreSQL Connected Successfully")
	return db
}
