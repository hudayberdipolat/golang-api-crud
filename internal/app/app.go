package app

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewApp(dependencies *AppDependencies) (httpServer *fiber.App) {
	httpServer = fiber.New(fiber.Config{
		AppName:                 dependencies.AppConfig.ServerConfig.AppName,
		BodyLimit:               35 * 1024 * 1024,
		EnableTrustedProxyCheck: true,
		ServerHeader:            dependencies.AppConfig.ServerConfig.AppHeader,
		WriteTimeout:            3 * time.Minute,
		ReadTimeout:             3 * time.Minute,
	})
	return httpServer
}
