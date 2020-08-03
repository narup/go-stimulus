package view

import "github.com/gofiber/fiber"


func Home(ctx *fiber.Ctx) {
	m := New("index", "Starter project for golang and stimulus js")
	render(ctx,m)
}

func Login(ctx *fiber.Ctx) {
	m := New("login", "Log In")
	render(ctx,m)
}
