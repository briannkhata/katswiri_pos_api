// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/briannkhata/katswiri_pos_api/database"
// 	"github.com/briannkhata/katswiri_pos_api/router"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/joho/godotenv"
// )

// var (
// 	DBConn *gorm.DB
// )

// func init() {

// 	if err := godotenv.Load(".env"); err != nil {
// 		log.Fatal("Error in loading .env file")
// 	}

// 	database.InitDB()
// }
// func main() {

// 	app := fiber.New(fiber.Config{
// 		CaseSensitive: true,
// 		AppName:       "Katswiri API v1.0.1",
// 	})

// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "*",
// 		AllowMethods: "PUT, PATCH, GET, POST, DELETE",
// 		AllowHeaders: "Origin, Auth-token, token, Content-type",
// 	}))

// 	app.Use(logger.New())
// 	router.SetUpRoutes(app)
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "10000"
// 	}

// 	// Start Fiber server and listen on all interfaces
// 	// if err := app.Listen("0.0.0.0:" + port); err != nil {
// 	// 	log.Fatal("Error starting server:", err)
// 	// }

// 	// Start the server
// 	// if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
// 	// 	log.Fatal("Error starting server:", err)
// 	// }
// 	app.Listen(":" + port)
// }

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

	// Dynamically fetch the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port for local development
	}

	// Start the server on the assigned port
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
