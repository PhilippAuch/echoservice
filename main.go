package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	m := make(map[string]float64)
	m["init"] = 0.0
	m["second"] = 42

	app.Get("/data/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER /data/ GET arguments: key=" + key)
		for key, element := range m {
			fmt.Println(key, []byte(key), element)
		}
		fmt.Println(time.Now().Format(time.RFC3339)+" map:", m)
		for key, element := range m {
			fmt.Println(key, []byte(key), element)
		}
		fmt.Println(time.Now().Format(time.RFC3339) + " EXIT /data/ GET")
		if len(key) > 0 {
			return c.JSON(m[key])
		}
		return c.JSON(m)
	})

	app.Get("/metrics", func(c *fiber.Ctx) error {
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER /metrics GET")
		for key, element := range m {
			fmt.Println(key, []byte(key), element)
		}

		result := "echoservicestatus 1"

		for key, element := range m {
			result = result + "\nechoservicedata{key=\"" + key + "\"} " + fmt.Sprintf("%f", element)
		}

		for key, element := range m {
			fmt.Println(key, []byte(key), element)
		}
		fmt.Println(time.Now().Format(time.RFC3339) + " EXIT /metrics GET")
		return c.SendString(result)
	})

	app.Post("/data/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		fmt.Println(time.Now().Format(time.RFC3339), " ENTER / POST arguments: key=", key)
		for key, element := range m {
			fmt.Println(key, []byte(key), element)
		}
		fmt.Println(time.Now().Format(time.RFC3339), " map:", m)
		a, b := strconv.ParseFloat(string(c.Body()), 64)
		fmt.Println(a, b)
		m[key] = a
		fmt.Println(time.Now().Format(time.RFC3339), " map:", m)
		for key, element := range m {
			fmt.Println(key, []byte(key), element)
		}
		fmt.Println(time.Now().Format(time.RFC3339), " EXIT / POST")
		return c.SendString("Done")
	})

	app.Listen(":8080")
}
