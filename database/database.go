// package database

// import (
// 	"log"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// var (
// 	DBConn *gorm.DB
// )

// func InitDB() {
// 	var err error
// 	dsn := "root@tcp(localhost:3306)/hokala?charset=utf8mb4&parseTime=True&loc=Local"
// 	DBConn, err = gorm.Open("mysql", dsn)

// 	if err != nil {
// 		log.Fatalf("Failed to connect to MySQL database: %v", err)
// 	}

// 	DBConn.LogMode(true)
// 	log.Println("Connected to MySQL database successfully.")
// }

package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var err error
	dsn := "root@tcp(localhost:3306)/hokala?charset=utf8mb4&parseTime=True&loc=Local"

	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}

	sqlDB, err := DBConn.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB instance: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Connected to MySQL database successfully.")
}
