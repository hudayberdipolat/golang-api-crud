package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/golang-api-crud/internal/domain/post/constructor"
)

func Routes(app *fiber.App) {
	apiRoute := app.Group("/api")
	// post route
	postRoute := apiRoute.Group("posts")
	postRoute.Get("", constructor.PostHandler.GetAll)
	postRoute.Get("/:postID", constructor.PostHandler.GetOne)
	postRoute.Post("/create", constructor.PostHandler.Create)
	postRoute.Put("/:postID/update", constructor.PostHandler.Update)
	postRoute.Delete("/:postID/delete", constructor.PostHandler.Delete)
}
