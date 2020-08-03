package view

import (
	"github.com/gofiber/fiber"
	"log"
)

type Model struct {
	Name string
	Title string
	Data map[string]interface{}
}

func New(name, title string) *Model {
	return &Model{Name: name}
}

func SetupRoutes(app *fiber.App) {
	app.Static("/static", "../webapp")

	app.Get("/", Home)
	app.Get("/login", Login)
}

func render(ctx *fiber.Ctx, m *Model) {
	err:=ctx.Render(m.Name, fiber.Map{"View":m},  "layouts/main")
	if err != nil {
		log.Printf("[ERROR][VIEW] error rendering a view %s. Error: %s", m.Name, err)
	}
}