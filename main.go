package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{ //frontend can get cookie
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")
}
