package server

import (
	"fmt"
	"log"
	"neon-api/server/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// StartServer inicializa el servidor Fiber
func StartServer(conn *pgx.Conn) { //requerimos la conexion con la base de datos "conn" para poder hacer solicitudes a esta misma.
	// Crear una nueva instancia de Fiber
	app := fiber.New()

	// SetupRoutes maneja las peticiones del cliente
	routes.SetupRoutes(app, conn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		fmt.Println("Listening in http://localhost:3000/...")
	}

	// Iniciar el servidor
	log.Fatal(app.Listen(":" + port))
}
