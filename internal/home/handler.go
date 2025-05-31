package home

import (
	"bytes"
	"html/template"

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
	tmpl := template.Must(template.ParseFiles("./html/page.html"))
	data := struct{ Count int }{Count: 5}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template compile error")
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML) // Задаем content-type в заголовке ответа
	return c.Send(tpl.Bytes())

}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Msg("Информация")

	return c.SendString("Hello")
}
