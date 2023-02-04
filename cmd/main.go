package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"myfiber-api/pkg/common/config"
	"myfiber-api/pkg/common/db"
	"myfiber-api/pkg/products"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	products.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
