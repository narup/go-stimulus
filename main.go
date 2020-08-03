package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/template/html"
	"github.com/narup/go-stimulus/internal/view"
	"net/http"
)

func main() {
	// We also support the http.FileSystem interface
	// See examples below to load templates from embedded files
	engine := html.NewFileSystem(http.Dir("../webapp/html"), ".html")
	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false
	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false

	// Layout defines the variable name that is used to yield templates within layouts
	engine.Layout("content") // Optional. Default: "embed"
	engine.AddFunc("content", func() error {
		return fmt.Errorf("layout called unexpectedly.")
	})
	// Delims sets the action delimiters to the specified strings
	engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	// After you created your engine, you can pass it to Fiber's Views Engine
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})
	app.Use(middleware.Recover())

	//setup routes
	view.SetupRoutes(app)

	app.Listen(3000)
}