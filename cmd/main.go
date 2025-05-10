package main

import (
	"develop/go-fiber/config"
	"develop/go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogCofig()

	log.SetLevel(log.Level(logConfig.Level))

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen("localhost:3000")
}
