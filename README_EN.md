# Golang CRUD Project (Neon-API)

This project is a RESTful API (CRUD) developed in **Go** using the **Fiber** framework and connected to a **PostgreSQL** database hosted on **Neon**.

## 🚀 Technologies

*   **Language**: [Go](https://go.dev/)
*   **Framework**: [Fiber](https://gofiber.io/) (Similar to Express.js for Node.js)
*   **Database**: [PostgreSQL](https://www.postgresql.org/) (Serverless via [Neon](https://neon.tech/))
*   **SQL Driver**: [pgx](https://github.com/jackc/pgx)
*   **Hot Reload**: [Air](https://github.com/cosmtrek/air)

## 🛠️ Prerequisites

*   **Go** installed (version 1.18 or higher).
*   A [Neon](https://neon.tech/) account and a database connection string.
*   (Optional) **Air** installed for hot reloading during development.

## ⚙️ Installation and Setup

1.  **Clone the repository:**
    ```bash
    git clone <REPOSITORY_URL>
    cd neon-api
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Configure environment variables:**
    Create a `.env` file in the project root and add your Neon connection string:
    ```env
    DATABASE_URL="postgres://<user>:<password>@<host>/<dbname>?sslmode=require"
    ```

## ▶️ Execution

### Development Mode (with Hot Reload)
If you have `air` installed, simply run:
```bash
air
```
This will automatically restart the server when changes are saved.

### Normal Mode
```bash
go run main.go
```

The server will start by default on port `3000`.

## 📚 API Documentation

### Users

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/users` | Retrieves all users. |
| `POST` | `/users` | Creates a new user. |
| `GET` | `/users/:id` | Finds a user by their `user_id`. |
| `DELETE` | `/users/:id` | Deletes a user by their `user_id`. |
| `GET` | `/users/last` | Retrieves the last registered user. |

### Utilities / Test

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | Welcome message (Health check). |
| `GET` | `/db` | DB route test message. |

### Testing Routes (Automatic)

These routes are designed for quick and automatic database testing.

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/db/post/test` | Automatically creates 3 test records with "is_test" set to true. |
| `GET` | `/db/get/test` | Shows records with "is_test" set to true. |
| `GET` | `/db/delete/test` | Deletes records with "is_test" set to true. |
| `GET` | `/db/drop` | **DANGER**: Deletes the entire `users` table. (The table is automatically recreated if it doesn't exist when the server starts). |

## 🧪 Usage Examples (cURL)

**Create a user:**
```bash
curl -X POST http://localhost:3000/users \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "john@example.com", "is_test": true}'

Returns:
```json
{
    "message": "User created successfully",
    "user created": {
        "record": 111,
        "user_id": "5c07f2f480",
        "name": "John Doe",
        "email": "john@example.com",
        "is_test": true
    }
}
```

**Find a user:**
```bash
curl http://localhost:3000/users/5c07f2f480

Returns:
```json
{
    "menssage": "User successfully found",
    "user": {
        "record": 111,
        "user_id": "5c07f2f480",
        "name": "John Doe",
        "email": "john@example.com",
        "is_test": true
    }
}
```

**Delete a user:**
```bash
curl -X DELETE http://localhost:3000/users/5c07f2f480

Returns:
```json
{
    "message": "User deleted successfully",
    "user deleted": {
        "record": 111,
        "user_id": "5c07f2f480",
        "name": "John Doe",
        "email": "john@example.com",
        "is_test": true
    }
}
```
**"is_test" is a boolean field that indicates if the user is a test user or not.**
