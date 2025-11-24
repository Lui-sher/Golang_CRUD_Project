# Golang CRUD Project (Neon-API)

Este proyecto es una API RESTful (CRUD) desarrollada en **Go** utilizando el framework **Fiber** y conectada a una base de datos **PostgreSQL** alojada en **Neon**.

## 🚀 Tecnologías

*   **Lenguaje**: [Go](https://go.dev/)
*   **Framework**: [Fiber](https://gofiber.io/) (Similar a Express.js para Node.js)
*   **Base de Datos**: [PostgreSQL](https://www.postgresql.org/) (Serverless via [Neon](https://neon.tech/))
*   **Driver SQL**: [pgx](https://github.com/jackc/pgx)
*   **Hot Reload**: [Air](https://github.com/cosmtrek/air)

## 🛠️ Requisitos Previos

*   Tener instalado **Go** (versión 1.18 o superior).
*   Una cuenta en [Neon](https://neon.tech/) y una cadena de conexión a la base de datos.
*   (Opcional) Tener instalado **Air** para recarga en caliente durante el desarrollo.

## ⚙️ Instalación y Configuración

1.  **Clonar el repositorio:**
    ```bash
    git clone <URL_DEL_REPOSITORIO>
    cd neon-api
    ```

2.  **Instalar dependencias:**
    ```bash
    go mod tidy
    ```

3.  **Configurar variables de entorno:**
    Crea un archivo `.env` en la raíz del proyecto y añade tu cadena de conexión de Neon:
    ```env
    DATABASE_URL="postgres://<user>:<password>@<host>/<dbname>?sslmode=require"
    ```

## ▶️ Ejecución

### Modo Desarrollo (con Hot Reload)
Si tienes `air` instalado, simplemente ejecuta:
```bash
air
```
Esto reiniciará el servidor automáticamente al guardar cambios.

### Modo Normal
```bash
go run main.go
```

El servidor iniciará por defecto en el puerto `3000`.

## 📚 Documentación de la API

### Usuarios

| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `GET` | `/users` | Obtiene todos los usuarios. |
| `POST` | `/users` | Crea un nuevo usuario. |
| `GET` | `/users/:id` | Busca un usuario por su `user_id`. |
| `DELETE` | `/users/:id` | Elimina un usuario por su `user_id`. |
| `GET` | `/users/last` | Obtiene el último usuario registrado. |

### Utilidades / Test

| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `GET` | `/` | Mensaje de bienvenida (Health check). |
| `GET` | `/db` | Mensaje de prueba de ruta DB. |

### Rutas de Testing (Automáticas)

Estas rutas están diseñadas para pruebas rápidas y automáticas de la base de datos.

| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `GET` | `/db/post/test` | Crea automáticamente 3 registros de prueba con la casilla "is_test" en true. |
| `GET` | `/db/get/test` | Muestra los registros con la casilla "is_test" en true. |
| `GET` | `/db/delete/test` | Elimina los registros con la casilla "is_test" en true. |
| `GET` | `/db/drop` | **PELIGRO**: Elimina la tabla `users` completa. (la tabla si no existe se crea automaticamente al iniciar el servidor) |

## 🧪 Ejemplos de Uso (cURL)

**Crear un usuario:**
```bash
curl -X POST http://localhost:3000/users \
-H "Content-Type: application/json" \
-d '{"name": "Juan Perez", "email": "juan@example.com", "is_test": true}'

Retorna:
```json
{
    "message": "User created successfully",
    "user created": {
        "record": 111,
        "user_id": "5c07f2f480",
        "name": "Juan Perez",
        "email": "juan@example.com",
        "is_test": true
    }
}
```

**Buscar un usuario:**
```bash
curl http://localhost:3000/users/5c07f2f480

Retorna:
```json
{
    "menssage": "User successfully found",
    "user": {
        "record": 111,
        "user_id": "5c07f2f480",
        "name": "Juan Perez",
        "email": "juan@example.com",
        "is_test": true
    }
}
```

**Eliminar un usuario:**
```bash
curl -X DELETE http://localhost:3000/users/5c07f2f480

Retorna:
```json
{
    "message": "User deleted successfully",
    "user deleted": {
        "record": 111,
        "user_id": "5c07f2f480",
        "name": "Juan Perez",
        "email": "juan@example.com",
        "is_test": true
    }
}
```
**"is_test" es un campo booleano que indica si el usuario es un usuario de prueba o no.**
