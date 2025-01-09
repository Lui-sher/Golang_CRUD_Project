package db

import (
	"context"
	"fmt"
	"neon-api/variables"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// Pasos para testear la base de datos
func PostTest(c *fiber.Ctx, conn *pgx.Conn) error {

	for i := 0; i < 3; i++ {
		variables.Count = i
		CreateUser(c, conn)
	}

	FetchDataTest(conn)
	return nil
}

// FetchDataTest recupera los datos de la tabla y los imprime.
func FetchDataTest(conn *pgx.Conn) error {
	query := `SELECT * FROM users`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error consultando datos: %w", err)
	}
	defer rows.Close()

	fmt.Printf("%6s |%10s   |%13s         |%17s\n", "Record", "User_Id", "Name", "Email")
	fmt.Println("-------+-------------+----------------------+----------------------------")
	for rows.Next() {
		var user variables.User
		if err := rows.Scan(&user.Record, &user.User_Id, &user.Name, &user.Email); err != nil {
			return fmt.Errorf("error escaneando fila: %w", err)
		}
		fmt.Printf("%4d   | %11s | %20s | %27s \n", user.Record, user.User_Id, user.Name, user.Email)
	}
	return nil
}
