package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// ConnectDB establece la conexión con la base de datos y la devuelve.
func ConnectDB() (*pgx.Conn, error) {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtener cadena de conexión
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL no está configurada en el archivo .env")
	}

	// Establecer conexión con la base de datos
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	fmt.Println("Conexión exitosa con la base de datos")

	// Eliminamos la tabla "users" si existe para dejar todo en cero
	state := false
	if state {
		if err := DropTable(conn); err != nil {
			panic(err)
		}
	}

	// Sincronizacion con la tabla 'users'
	if err := SetupDatabase(conn); err != nil {
		panic(err)
	}

	// Verificamos los datos ingresados
	// if err := FetchAllData(conn); err != nil {
	// 	panic(err)
	// }

	return conn, nil
}

// CloseDB cierra la conexión global a la base de datos.
func CloseDB(conn *pgx.Conn) {
	if conn != nil {
		conn.Close(context.Background())
		fmt.Println("Conexión a la base de datos cerrada")
	}
}
