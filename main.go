package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	m := make(map[string]float64)
	m["init"] = 0.0
	m["second"] = 42

	app.Get("/data/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER /data/ GET arguments: key=" + key)
		// fmt.Println(time.Now().Format(time.RFC3339)+" map:", m)
		// fmt.Println(time.Now().Format(time.RFC3339) + " EXIT /data/ GET")
		if len(key) > 0 {
			return c.JSON(m[key])
		}
		return c.JSON(m)
	})

	app.Get("/metrics", func(c *fiber.Ctx) error {
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER /metrics GET")

		result := "echoservicestatus 1"
		for key, element := range m {
			result = result + "\nechoservicedata{key=\"" + key + "\"} " + fmt.Sprintf("%f", element)
		}

		// fmt.Println(time.Now().Format(time.RFC3339) + " EXIT /metrics GET")
		return c.SendString(result)
	})

	app.Post("/data/:key?", func(c *fiber.Ctx) error {
		key := string(c.Params("key"))
		fmt.Println(time.Now().Format(time.RFC3339), " ENTER / POST arguments: key=", key)
		m[key], _ = strconv.ParseFloat(string(c.Body()), 64)
		// fmt.Println(time.Now().Format(time.RFC3339), " map:", m)
		// fmt.Println(time.Now().Format(time.RFC3339), " EXIT / POST arguments: key=", key)
		return c.JSON(m)
	})

	app.Listen(":8080")
}
