# InventoryManagementAPI

A minimal Go application demonstrating an Inventory Management API with:
- Gin web framework
- PostgreSQL via GORM
- Rate limiting
- Basic CRUD operations
- Pagination, filtering, and sorting

## Repository Structure

```
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
2. **Clone this repo**:
   ```bash
   git clone https://github.com/<YourUsername>/InventoryManagementAPI.git
   cd InventoryManagementAPI
   ```
3. **Initialize dependencies**:
   ```bash
   go mod tidy
   ```
4. **Set your DSN** (optional; otherwise a default DSN is used in code):
   ```bash
   export DSN="postgres://postgres:postgres@localhost:5432/inventory?sslmode=disable"
   ```
5. **Run**:
   ```bash
   go run ./cmd/server/main.go
   ```
6. **Test** the API at `http://localhost:8080`.

## Creating a GitHub Codespace

1. Push this code to a GitHub repository named `InventoryManagementAPI`.
2. On GitHub, click the green "Code" button and select "Create codespace on main."
3. In the Codespace, open a terminal and run:
   ```bash
   go mod tidy
   go run ./cmd/server/main.go
   ```
4. Expose or forward port `8080` in your Codespace to access the running server from your browser.

## Contact

This repository is maintained by **Moamen Abdelkawy**. For questions, suggestions, or collaboration inquiries, please reach out via email at [moamen.abdelkawy@outlook.com](mailto:moamen.abdelkawy@outlook.com).

## License

This project is licensed under the [MIT License](LICENSE).