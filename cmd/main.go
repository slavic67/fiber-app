package main

import (
	"develop/go-fiber/config"
	"develop/go-fiber/internal/home"
	"develop/go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogCofig()
	customLoggger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLoggger,
	}))
	app.Use(recover.New())

	home.NewHandler(app, customLoggger)
	app.Listen("localhost:3000")
}
