//https://www.youtube.com/watch?v=3JtZqqrJFmM&list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9&index=17

package main

import (
	"fmt"
	"ims-server/routes"

	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1/create/keyvalue", routes.CreateKeyValue)
	app.Delete("/api/v1/delete/keyvalue", routes.DeleteKeyValue)
	app.Get("/api/v1/find/keyvalue", routes.FindKeyValue)
	// app.Post("/api/v1/lead", lead.NewLead)
	// app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {
	// for removing all data remove directory /var/lib/etcd/default   and restart etcd service

	app := fiber.New()
	//initDatabase()
	setupRoutes(app)
	err := app.Listen(3001)

	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
	// defer database.DBConn.Close()
}
