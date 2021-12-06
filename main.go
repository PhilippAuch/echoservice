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

	app.Get("/hello/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER /hello GET arguments: name=" + name)
		if len(name) <= 0 {
			name = "World"
		}

		fmt.Println(time.Now().Format(time.RFC3339) + " EXIT  /hello GET arguments: name=" + name)
		return c.SendString("Hello " + name)
	})

	app.Get("/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		fmt.Println(time.Now().Format(time.RFC3339) + " ENTER / GET arguments: key=" + key)

		fmt.Println(time.Now().Format(time.RFC3339)+" map:", m)

		value := ""

		if len(key) > 0 {
			valueString, _ := json.Marshal(m[key])
			value = string(valueString)
		} else {
			valueString, _ := json.Marshal(m)
			value = string(valueString)
		}

		fmt.Println(time.Now().Format(time.RFC3339)+" EXIT  / GET arguments: key="+key+", value=", value)
		return c.SendString(value)
	})

	app.Listen(":8080")
}
