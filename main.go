package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/robjullian/redimail/controller"
	"github.com/robjullian/redimail/database"
)

// Load env values
func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Unable to load env values: %v\n", err)
	} else {
		log.Println("Loaded env values successfully")
	}

	if os.Getenv("DOMAIN") == "" ||
		os.Getenv("SERVER_HOST") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_PATH") == "" ||
		os.Getenv("DB_FILE") == "" {
		log.Fatalln("Please set these environment variables: \n - DOMAIN\n - SERVER_HOST\n - SERVER_PORT\n - DB_PATH\n - DB_FILE")
	}
}

func main() {
	app := fiber.New()

	// Create the connection to the database database
	if err := database.Connect(); err != nil {
		log.Printf("Unable to connect the database: %v\n", err.Error())
		panic(err)
	} else {
		log.Println("Successfully connected to the database")
	}

	app.Get("/:id", controller.GetEmail)
	app.Post("/add", controller.AddEmail)

	// Create listener
	listner := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	err := app.Listen(listner)
	if err != nil {
		log.Println("Unable to start server on [" + listner + "]")
		panic(err)
	} else {
		log.Println("Listening on [" + listner + "]")
	}
}
