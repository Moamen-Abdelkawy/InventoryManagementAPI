# **InventoryManagementAPI**

A minimal Go application demonstrating an Inventory Management API with:
- Gin web framework
- PostgreSQL via GORM
- Rate limiting
- Basic CRUD operations
- Pagination, filtering, and sorting

## Repository Structure

```plaintext
InventoryManagementAPI/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── database/
│   │   └── database.go
│   ├── handlers/
│   │   └── items.go
│   ├── middleware/
│   │   └── ratelimit.go
│   ├── models/
│   │   └── item.go
│   └── router/
│       └── routes.go
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Running Locally

1. **Install Go 1.18+** and ensure your `$GOPATH` or modules are set up.

2. Clone this repo:

   ```bash
   git clone https://github.com/<YourUsername>/InventoryManagementAPI.git
   cd InventoryManagementAPI
   ```

3. Initialize dependencies:

   ```bash
   go mod tidy
   ```

4. Set your DSN (optional; otherwise a default DSN is used in code):

   ```bash
   export DSN="postgres://postgres:postgres@localhost:5432/inventory?sslmode=disable"
   ```

5. Run:

   ```bash
   go run ./cmd/server/main.go
   ```

   > **Note**: The server will keep running and listen on `http://localhost:8080` until stopped (e.g., via `Ctrl+C`).

## Creating a GitHub Codespace

1. Push this code to a GitHub repository named `InventoryManagementAPI`.

2. On GitHub, click the green "Code" button and select "Create codespace on main."

3. In the Codespace, open a terminal and run:

   ```bash
   go mod tidy
   go run ./cmd/server/main.go
   ```

4. Forward port `8080` in your Codespace to access the running server from your browser.

## Database Options in a Codespace

1. **Docker-based Postgres**

   ```bash
   docker run -d --name postgres -p 5432:5432 \
     -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres postgres:15
   export DSN="postgres://postgres:postgres@127.0.0.1:5432/inventory?sslmode=disable"
   go run ./cmd/server/main.go
   ```

2. **Hosted Postgres (e.g., ElephantSQL)**

   - Sign up for a free ElephantSQL instance.

   - Copy the connection URL (DSN), then:

     ```bash
     export DSN="<YourElephantSQLConnectionString>"
     go run ./cmd/server/main.go
     ```

## Troubleshooting

- If you see `connection refused` errors, confirm that PostgreSQL is running (in Docker or a hosted provider).
- If `go` commands fail, ensure Go is installed and that you’re in the repo root with `go.mod`.

## Contact

This repository is maintained by **Moamen Abdelkawy**. For questions, suggestions, or collaboration inquiries, please reach out via email at [moamen.abdelkawy@outlook.com](mailto:moamen.abdelkawy@outlook.com).

## License

This project is licensed under the [MIT License](https://chatgpt.com/c/LICENSE).
