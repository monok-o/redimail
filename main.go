package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/robjullian/fff-mail-link/database"
	"github.com/robjullian/fff-mail-link/models"
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

	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var mail models.Mail

		database.DB.Model(&models.Mail{}).Where("id = ?", id).Find(&mail)

		if mail.To != "" {
			return c.Redirect(fmt.Sprintf("mailto:%s?cc=%s&bcc=%s&subject=%s&body=%s", mail.To, mail.Cc, mail.Bcc, mail.Subject, mail.Content))
		} else {
			return c.SendString("Not found")
		}
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Unable to process JSON request",
			})
		}

		u, err := uuid.NewV4()

		mail := models.Mail{
			Id:      u.String(),
			To:      data["to"],
			Cc:      data["cc"],
			Bcc:     data["bcc"],
			Subject: data["subject"],
			Content: data["content"],
		}

		err = database.DB.Create(mail).Error
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
				"message": "Database error",
				"error":   err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Mail created",
			"link":    fmt.Sprintf("%s/%s", os.Getenv("DOMAIN"), u.String()),
		})
	})

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
