package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func StartServer(port int) error {
	app := fiber.New()
	app.Post("/todo/add", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return nil
	})
	err := app.Listen(fmt.Sprintf(":%d", port))
	return err
}
