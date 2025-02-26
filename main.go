package main

import (
	"log"
	"os"

	"github.com/briannkhata/katswiri_pos_api/database"
	"github.com/briannkhata/katswiri_pos_api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	DBConn *gorm.DB
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file")
	}

	database.InitDB()
}
func main() {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "Katswiri API v1.0.1",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "PUT, PATCH, GET, POST, DELETE",
		AllowHeaders: "Origin, Auth-token, token, Content-type",
	}))

	app.Use(logger.New())
	router.SetUpRoutes(app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(port)

}
