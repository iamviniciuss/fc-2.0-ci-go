package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	fmt.Println("Starting app 009")
	go running()

	app := fiber.New()

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")
		c.Set("Access-Control-Allow-Origin", "*")

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Request /")
		c.Status(200)
		return c.Send([]byte("<h1>Hello 222</h1>"))
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		fmt.Println("health /")
		c.Status(200)
		return c.Send([]byte("<h1>health 2</h1>"))
	})

	log.Fatal(app.Listen(":80"))
}

func running() {
	for {
		fmt.Println("Running")
		time.Sleep(time.Second * 10)
	}
}
