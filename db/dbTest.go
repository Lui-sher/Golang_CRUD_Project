package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"neon-api/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// ---------------- Funciones para testear la base de datos -------------------

// PostTest crea tres nuevos registros de usuarios automaticamente en la base de datos
func PostTest(c *fiber.Ctx, conn *pgx.Conn) error {
	client := &http.Client{}
	var results []map[string]any

	for i := range 3 {
		body := fmt.Sprintf(`{
			"name": "Usuario de prueba %d",
			"email": "usuarioDePrueba%d@prueba.com",
			"is_test": true
		}`, i+1, i+1)

		req, _ := http.NewRequest("POST", "http://localhost:3000/db/create/user", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error enviando request:", err)
			continue
		}
		//fmt.Println(resp.Body)
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error leyendo datos: %w", err)
		}

		results = append(results, fiber.Map{
			"request":  i + 1,
			"response": json.RawMessage(bodyBytes),
		})
	}

	if err := FetchAllData(conn); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": results,
	})
}

// GetTest: consulta y muestra todos los registros creados por la funcion PostTest
func GetTest(conn *pgx.Conn) error { // In Process...
	query := `SELECT * FROM users WHERE is_test = TRUE`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error consultando datos: %w", err)
	}
	defer rows.Close()

	if err := ShowRowsInTerminal(rows); err != nil {
		return fmt.Errorf("algo salió mal: %w", err)
	}

	return nil
}

// DeleteTest: elimina todos los registros creados por la funcion PostTest
func DeleteTest(conn *pgx.Conn) error {
	query := `DELETE FROM users WHERE is_test = TRUE`

	// Ejecutar la consulta
	cmdTag, err := conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error eliminando datos: %w", err)
	}

	// Mostrar cuántas filas se eliminaron
	fmt.Printf("Se eliminaron %d registros de prueba de la tabla 'users'.\n", cmdTag.RowsAffected())

	return nil
}

// Mostrar la tabla por consola
func ShowRowsInTerminal(rows pgx.Rows) error {

	fmt.Printf("%6s |%10s   |%13s         |%17s\n", "Record", "User_Id", "Name", "Email")
	fmt.Println("-------+-------------+----------------------+----------------------------")
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Record, &user.User_Id, &user.Name, &user.Email, &user.Is_test); err != nil {
			return fmt.Errorf("error escaneando fila: %w", err)
		}
		fmt.Printf("%4d   | %11s | %20s | %27s \n", user.Record, user.User_Id, user.Name, user.Email)
	}
	return nil
}

// FetchDataTest recupera todos los datos de la tabla y los muestra por consola.
func FetchAllData(conn *pgx.Conn) error {
	query := `SELECT * FROM users`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error consultando datos: %w", err)
	}
	defer rows.Close()

	if err := ShowRowsInTerminal(rows); err != nil {
		return fmt.Errorf("algo salió mal: %w", err)
	}

	return nil
}
