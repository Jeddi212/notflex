package controllers

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect | Database configuration - GORM configuration
func Connect() *gorm.DB {
	// Set custom logger for GORM Framework
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level  // Ignore ErrRecordNotFound error for logger
			Colorful:      false,         // Disable color
		},
	)

	// Set database information
	dsn := "root@tcp(localhost:3306)/db_notflex?charset=utf8mb4&parseTime=true&loc=Asia%2FJakarta"

	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
