package main

import (
	"neon-api/db"
	"neon-api/server"
)

func main() {
	//Conexeion con la base de datos
	conn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Nos aseguramos de cerrar la conexión a la base de datos al final
	defer db.CloseDB(conn)

	//Levantamos el servidor, para la comunicacion con el cliente
	server.StartServer(conn)
}
