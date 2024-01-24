package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/golang-api-crud/pkg/config"
)

func SetStaticRoute(app *fiber.App, config config.Config) {
	app.Static("/public", config.PublicPath)
}
