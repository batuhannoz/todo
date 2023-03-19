package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"net/http"
)

func main() {
	err := StartServer(3000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Provider Service Listening :3000")
}

func StartServer(port int) error {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/todo/add", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return nil
	})
	err := app.Listen(fmt.Sprintf(":%d", port))
	return err
}
