package db

import (
	"context"
	"fmt"
	"neon-api/variables"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// ---------------- Pasos para testear la base de datos -------------------

// PostTest crea tres nuevos registros de usuarios en la base de datos
func PostTest(c *fiber.Ctx, conn *pgx.Conn) error {

	for i := 0; i < 3; i++ {
		c.Request().SetBody([]byte(fmt.Sprintf(`{
		"name": "Usuario de prueba %d",
		"email": "usuarioDePrueba%d@prueba.com",
		"is_test": true
		}`, i+1, i+1)))
		if err := CreateUser(c, conn); err != nil {
			return err
		}
	}
	if err := FetchAllData(conn); err != nil {
		return err
	}
	return nil
}

// GetTest, consulta y muestra todos los registros creados por la funcion PostTest
func GetTest(conn *pgx.Conn) error {

	return nil
}

func ShowRowsInTerminal(rows pgx.Rows) error {

	fmt.Printf("%6s |%10s   |%13s         |%17s\n", "Record", "User_Id", "Name", "Email")
	fmt.Println("-------+-------------+----------------------+----------------------------")
	for rows.Next() {
		var user variables.User
		if err := rows.Scan(&user.Record, &user.User_Id, &user.Name, &user.Email, &user.Is_test); err != nil {
			return fmt.Errorf("error escaneando fila: %w", err)
		}
		fmt.Printf("%4d   | %11s | %20s | %27s \n", user.Record, user.User_Id, user.Name, user.Email)
	}
	return nil
}

// FetchDataTest recupera los datos de la tabla y los muestra por consola.
func FetchAllData(conn *pgx.Conn) error {
	query := `SELECT * FROM users`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error consultando datos: %w", err)
	}
	defer rows.Close()

	if err := ShowRowsInTerminal(rows); err != nil {
		return fmt.Errorf("Algo salió mal: %w", err)
	}

	return nil
}
