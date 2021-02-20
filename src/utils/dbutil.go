package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Gorm *gorm.DB
func InitDB(){
	Gorm=gormDB()
}
func gormDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	dsn:="root:root@tcp(localhost:3306)/oauth?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{Logger:newLogger})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB,err:=db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetConnMaxLifetime(time.Second*30)
	return db
}