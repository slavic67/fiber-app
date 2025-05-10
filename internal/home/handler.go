package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().
		Bool("isAdmin", true).
		Str("email", "test@.ru").
		Msg("Информация")

	return c.SendString("Hello")

}
