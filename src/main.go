package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"

	"github.com/Vinicius-Santos-da-Silva/fc-2.0-ci-go/src/infra"
)

func main() {

	fmt.Println("Starting app 009")
	go running()

	infra.GetSecret()

	engine := html.New("./src/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

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
		return c.Send([]byte("<h1>Hello 34</h1>"))
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		fmt.Println("health /")
		c.Status(200)
		return c.Send([]byte("<h1>health 34</h1>"))
	})

	app.Get("/index", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("test1", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	log.Fatal(app.Listen(":80"))
}

func running() {
	for {
		fmt.Println("Running")
		time.Sleep(time.Second * 10)
	}
}
