package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()
	redis := rdInit()

	app.Get("/", func(c *fiber.Ctx) error {
		setKey(redis, "test_key", "value")
		message := getKey(redis, "test_key")

		return c.JSON(fiber.Map{
			"message": message,
		})
	})

	app.Listen(getPort())
}
