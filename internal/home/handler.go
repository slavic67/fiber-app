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
	api := h.router.Group("/api")
	api.Get("/", h.home) // end-point главной страницы
	api.Get("/error", h.error)
}

// Функция для end-point главной страницы
func (h *HomeHandler) home(c *fiber.Ctx) error {
	users := []User{
		{Id: 1, Name: "Anton"},
		{Id: 2, Name: "Vasia"},
	}
	names := []string{"Anton", "Vasia"}
	data := struct {
		Names []string
		Users []User
	}{Names: names, Users: users}
	return c.Render("page", data)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Msg("Информация")

	return c.SendString("Hello")
}
