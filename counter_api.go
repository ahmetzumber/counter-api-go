package main

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data int `json:"counter"`
}

var count int

func IncreaseHandler(c *fiber.Ctx) error {
	count++
	req := Response{Data: count}
	return c.Status(fiber.StatusOK).JSON(req)
}

func DecreaseHandler(c *fiber.Ctx) error {
	count -= 1
	req := Response{Data: count}
	return c.Status(fiber.StatusOK).JSON(req)
}

func DisplayHandler(c *fiber.Ctx) error {
	req := Response{Data: count}
	return c.Status(fiber.StatusOK).JSON(req)
}

func main() {
	// Start a new fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hello! Welcome Counter API.")
		return err
	})

	app.Get("/display", DisplayHandler)
	app.Get("/increase", IncreaseHandler)
	app.Get("/decrease", DecreaseHandler)

	// Listen on PORT 8080
	app.Listen(":8080")
}
