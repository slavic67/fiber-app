package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Msg("Информация")

	return c.SendString("Hello")

}
