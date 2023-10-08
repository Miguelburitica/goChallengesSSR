package main

import (
	"github.com/Miguelburitica/backend-project/go-server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
    // Initialize a new Fiber application
    app := fiber.New(fiber.Config{
        Views: html.New("./templates", ".html"),
    })

    // Serve static files
    app.Static("/", "./public")

    // Define a route
    app.Get("/", func(c *fiber.Ctx) error {
        // Render an HTML template
        return c.Render("index", fiber.Map{})
    })

    // todo challenge routes
    app.Get("/challenges/:day?/:withHeader?", func(c *fiber.Ctx) error {
        challengeDay := "day_1"
        
        day := c.Params("day")
        withHeader := c.Params("withHeader")
        
        if day != "" {
            challengeDay = day
        }
        if withHeader == "false" {
            return c.Render("challenges/" + challengeDay, fiber.Map{})
        }

        return c.Render("challenges/withHeader/" + challengeDay, fiber.Map{
            "currentChallengeNumber": day,
            "isChallengePage": true,
        })
    })

    // files routes
    routes.LoadFileRoutes(app)

    // todo user routes
    
    // Start the server
    app.Listen(":8080")
}
