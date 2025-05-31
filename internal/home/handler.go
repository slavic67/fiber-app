package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// Конфигурация обработчика
type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

// Конструктор обработчика
func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home) // end-point главной страницы
	api.Get("/error", h.error)
}

// Функция для end-point главной страницы
func (h *HomeHandler) home(c *fiber.Ctx) error {
	data := struct {
		Count   int
		IsAdmin bool
		CanUse  bool
	}{Count: 5, IsAdmin: true, CanUse: true}
	return c.Render("page", data)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Msg("Информация")

	return c.SendString("Hello")
}
