package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func Run() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":%s", port))

}
