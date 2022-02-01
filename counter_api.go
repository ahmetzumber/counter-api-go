package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CounterJson struct {
	counter int `json:counter`
}

var data CounterJson

func increaseHandler(c *fiber.Ctx) error {
	data.counter++
	return c.SendString(fmt.Sprint(data.counter))
}

func decreaseHandler(c *fiber.Ctx) error {
	data.counter -= 1
	return c.SendString(fmt.Sprint(data.counter))
}

func displayCounter(c *fiber.Ctx) error {
	var counterJson string = fmt.Sprintf(`{ "counter": %d }`, data.counter)
	return c.SendString(counterJson)
}

func main() {
	// Start a new fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	app.Get("/display", displayCounter)
	app.Get("/increase", increaseHandler)
	app.Get("/decrease", decreaseHandler)

	// Listen on PORT 3000
	app.Listen(":3000")
}
