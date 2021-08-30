package controller

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/robjullian/redimail/database"
	"github.com/robjullian/redimail/models"
)

func AddEmail(c *fiber.Ctx) error {
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
}
