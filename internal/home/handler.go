package home

import (
	"develop/go-fiber/pkg/tadapter"
	"develop/go-fiber/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// Конфигурация обработчика
type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type User struct {
	Id   int
	Name string
}

// Конструктор обработчика
func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}

	h.router.Get("/", h.home) // end-point главной страницы
	h.router.Get("/404", h.error)
}

// Функция для end-point главной страницы
func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main() // Создаем компонент на основе view
	return tadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Msg("Информация")

	return c.SendString("Hello")
}
