package main

import (
	"log"

	"goalculator/cmd"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ap := fiber.New()
	app := cmd.NewApp()

	ap.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	ap.Post("/api/calculate", func(c *fiber.Ctx) error {
		var body struct {
			A         float64 `json:"a"`
			B         float64 `json:"b"`
			Operation string  `json:"operation"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		result, err := app.Calculator.PerformOperation(body.A, body.B, body.Operation)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"result": result})
	})

	log.Fatal(ap.Listen(":3000"))
}
