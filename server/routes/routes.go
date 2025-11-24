package routes

import (
	"neon-api/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
)

// Creamos la funcion que manejará las peticiones del cliente, "conn" permite conectar con la base de datos
func SetupRoutes(app *fiber.App, conn *pgx.Conn) {

	//Middleware
	app.Use(logger.New())

	//------------------------- Peticiones GET-------------------------------------------------
	// Ruta "/"
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("¡Wellcome to Main Page!") })

	// Ruta de ejemplo para testear la base de datos
	app.Get("/db", func(c *fiber.Ctx) error { return c.SendString("Aquí irán las operaciones con la base de datos") })

	// Ruta para consultar el último registro ingresado en la tabla
	app.Get("/users/last", func(c *fiber.Ctx) error { return db.LastRecordHandler(c, conn) })

	// Ruta para localizar un usuario aportando su "user_id"
	app.Get("/users/:id", func(c *fiber.Ctx) error { return db.FindUserHandler(c, conn) })

	// Ruta para mostrar toda la tabla "Users" por consola
	app.Get("/users", func(c *fiber.Ctx) error { return db.FetchAllData(conn) })

	// ------------------------ Peticiones POST --------------------------------------------
	app.Post("/users", func(c *fiber.Ctx) error { return db.CreateUser(c, conn) })

	// -----------------------  Peticiones DELETE  ----------------------------------------
	app.Delete("/users/:id", func(c *fiber.Ctx) error { return db.DeleteUserHandler(c, conn) })

	// ------------------------ Testing Routes ---------------------------------------------
	// PostTest: Crear 3 nuevos registros automaticamente
	app.Get("/db/post/test", func(c *fiber.Ctx) error { return db.PostTest(c, conn) })

	// GetTest: Muestra todos los registro creados por el test anterior
	app.Get("/db/get/test", func(c *fiber.Ctx) error { return db.GetTest(conn) })

	// DeleteTest: Elimina todos los registros creados por PostTest
	app.Get("/db/delete/test", func(c *fiber.Ctx) error { return db.DeleteTest(conn) })

	// DropTable: elimina toda la tabla "users"
	app.Get("/db/drop", func(c *fiber.Ctx) error { return db.DropTable(conn) })
}
