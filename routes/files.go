package routes

import (
	"github.com/Miguelburitica/backend-project/go-server/controllers"
	"github.com/gofiber/fiber/v2"
)

func LoadFileRoutes(app *fiber.App) {
	app.Get("/upload-file", controllers.UploadFilesViewController)
	app.Post("/upload-file", controllers.UploadFileController)
}