package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	m := make(map[string]string)
	m["init"] = "init response"
	m["second"] = "5"

	app.Get("/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER / GET arguments: key=" + key)

		fmt.Println(time.Now().Format(time.RFC3339)+" map:", m)

		if len(key) > 0 {
			return c.JSON(m[key])
		}
		return c.JSON(m)
	})

	app.Post("/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER / POST arguments: key=" + key)

		value := ""
		if err := c.BodyParser(value); err != nil {
			errString, _ := json.Marshal(err)
			return c.Status(503).Send(errString)
		}
		m[key] = value

		fmt.Println(time.Now().Format(time.RFC3339)+" map:", m)

		return c.JSON(m)
	})

	app.Listen(":8080")
}
