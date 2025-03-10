package routes

import (
	"redock/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// TunnelRoutes func for describe group of private routes.
func TunnelRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Get("/tunnel/check_login", controllers.CheckUser)
	route.Post("/tunnel/login", controllers.TunnelLogin)
	route.Get("/tunnel/list", controllers.TunnelList)
	route.Post("/tunnel/delete", controllers.TunnelDelete)
	route.Post("/tunnel/add", controllers.TunnelAdd)
	route.Post("/tunnel/start", controllers.TunnelStart)
	route.Post("/tunnel/stop", controllers.TunnelStop)
}
