package main

import (
    "github.com/gofiber/fiber/v2"
    "fmt"
    "time"
)

func main() {
	app := fiber.New()

	app.Get("/hello/:name?", func(c *fiber.Ctx) error {
        name := c.Params("name")
        fmt.Println(time.Now().Format(time.RFC3339) + " ENTER / GET arguments: name=" + name)
        if len(name) <= 0 {
            name = "World"
        }

        fmt.Println(time.Now().Format(time.RFC3339) + " EXIT  / GET arguments: name=" + name)
        return c.SendString("Hello " + name)
	})

    app.Get("/Joke", func(c *fiber.Ctx) error {
        
        return c.SendString("Done")
    })

	app.Listen(":8080")
}