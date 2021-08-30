package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/robjullian/redimail/database"
	"github.com/robjullian/redimail/models"
)

func GetEmail(c *fiber.Ctx) error {
	id := c.Params("id")

	var mail models.Mail

	database.DB.Model(&models.Mail{}).Where("id = ?", id).Find(&mail)

	if mail.To != "" {
		return c.Redirect(fmt.Sprintf("mailto:%s?cc=%s&bcc=%s&subject=%s&body=%s", mail.To, mail.Cc, mail.Bcc, mail.Subject, mail.Content))
	} else {
		return c.SendString("Not found")
	}
}
