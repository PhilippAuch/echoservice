package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	m := make(map[string]string)
	m["init"] = "init response"
	m["second"] = "5"

	app.Get("/data/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER / GET arguments: key=" + key)
		fmt.Println(time.Now().Format(time.RFC3339)+" map:", m)

		if len(key) > 0 {
			return c.JSON(m[key])
		}
		return c.JSON(m)
	})

	app.Get("/metrics", func(c *fiber.Ctx) error {
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER /metrics/ GET")

		result := "echoservicestatus 1"
		return c.Send([]byte(result))
	})

	app.Post("/data/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER / POST arguments: key=" + key)
		m[key] = string(c.Body())
		fmt.Println(time.Now().Format(time.RFC3339)+" map:", m)
		return c.JSON(m)
	})

	app.Listen(":8080")
}
