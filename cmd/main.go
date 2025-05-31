package main

import (
	"develop/go-fiber/config"
	"develop/go-fiber/internal/home"
	"develop/go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {

	config.Init()                                // Загрузка файла .env
	config.NewDatabaseConfig()                   // Конфигуратор параметров доступа к базе данных
	logConfig := config.NewLogCofig()            // Параметры для логирования приложения
	customLoggger := logger.NewLogger(logConfig) // Настройка вида логирования приложения
	engine := html.New("./html", ".html")        //Загружаем шаблоны страниц

	app := fiber.New(fiber.Config{
		Views: engine,
	}) // Создание приложения
	app.Use(fiberzerolog.New(fiberzerolog.Config{ // Включение логирования на основе конфигурации
		Logger: customLoggger,
	}))
	app.Use(recover.New())              //Защищает приложения от паник в контроллерах
	home.NewHandler(app, customLoggger) // Контроллер для домашней страницы
	app.Listen("localhost:3000")        //Запуск приложения
}
