package db

import (
	"context"
	"fmt"
	"neon-api/variables"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// Eliminamos la tabla si exixte para dejar todo en cero
func DropTable(conn *pgx.Conn) error {
	// Verificar si la conexión está activa
	if conn == nil {
		return fmt.Errorf("la conexión a la base de datos no está inicializada")
	}

	query := `DROP TABLE IF EXISTS users;`
	if _, err := conn.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("error al intentar eliminar la tabla 'user': %v", err)
	}

	fmt.Println("La tabla 'users' ha sido eliminada exitosamente")
	return nil
}

// SetupDatabase configura la base de datos creando la tabla 'users' si no existe.
func SetupDatabase(conn *pgx.Conn) error {
	// Habilitar la extensión pgcrypto si no está habilitada (se usa para generar valores aleatorios)
	if _, err := conn.Exec(context.Background(), "CREATE EXTENSION IF NOT EXISTS pgcrypto;"); err != nil {
		return fmt.Errorf("error habilitando la extensión pgcrypto: %w", err)
	}

	query := `
		CREATE TABLE IF NOT EXISTS users(
			record SERIAL PRIMARY KEY,
			user_id VARCHAR(10) UNIQUE DEFAULT SUBSTRING(REPLACE(gen_random_uuid()::TEXT, '-', ''), 1, 10),
			name TEXT NOT NULL,
			email TEXT NOT NULL
	);`

	if _, err := conn.Exec(context.Background(), query); err != nil {
		panic(err)
	}
	fmt.Println("Tabla 'users' verificada o creada exitosamente")

	return nil
}

// LastRecord consulta el ultimo registro ingresado en la tabla 'users'
func LastRecord(conn *pgx.Conn) (*variables.User, error) {

	query := `SELECT * FROM users ORDER BY record DESC LIMIT 1;`
	row := conn.QueryRow(context.Background(), query)

	var user variables.User // Crear un puntero a una estructura de User

	// Escanear la fila obtenida
	err := row.Scan(&user.Record, &user.User_Id, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el último registro: %w", err)
	}

	return &user, nil
}

func LastRecordHandler(c *fiber.Ctx, conn *pgx.Conn) error {
	user, err := LastRecord(conn)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Responder con el registro encontrado
	return c.JSON(fiber.Map{
		"Record":  user.Record,
		"User_Id": user.User_Id,
		"name":    user.Name,
		"email":   user.Email,
	})
}

// Función Creadora de nuevos usuarios
func CreateUser(c *fiber.Ctx, conn *pgx.Conn) error {
	var user variables.User

	// ---------------- Testing ------------------
	if c.Path() == "/db/test" {

		user.Name = fmt.Sprintf("Usuario de prueba %d", variables.Count)
		user.Email = fmt.Sprintf("usuarioDePrueba%d@prueba.com", variables.Count)
		goto saltoTest
	}
	//--------------------------------------------

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	// Validamos los campos obligatorios
	if user.Name == "" || user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and Email are required fields",
		})
	}

saltoTest:
	query := `
        INSERT INTO users (name, email)
        VALUES ($1, $2)
		RETURNING user_id;
    `

	// Ejecutar la consulta POST
	if err := conn.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.User_Id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert user into the database",
		})
	}
	fmt.Println("Datos insertados exitosamente")

	// Responder al cliente con éxito
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Data inserted successfully",
		"user_id":    user.User_Id,
		"user_name":  user.Name,
		"user_email": user.Email,
	})
}

// Funcion para encontar a user
func FindUser(conn *pgx.Conn, userId string) (*variables.User, error) {

	query := "SELECT record, user_id, name, email FROM users WHERE user_id = $1;"
	row := conn.QueryRow(context.Background(), query, userId)

	var user variables.User
	err := row.Scan(&user.Record, &user.User_Id, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el usuario con user_id %s: %v", userId, err)
	}

	return &user, nil
}

func FindUserHandler(c *fiber.Ctx, conn *pgx.Conn) error {

	userId := c.Params("user_id")
	fmt.Println("Searching User_Id ", userId, "...")

	user, err := FindUser(conn, userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"Record":  user.Record,
		"User_Id": user.User_Id,
		"name":    user.Name,
		"email":   user.Email,
	})
}
