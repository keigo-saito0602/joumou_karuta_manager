# Joumou Karuta Manager

This is a backend server project built with Go (Echo + GORM), providing a RESTful API with Swagger documentation and database management via migrations.

---

## ✨ Features

- RESTful API for managing Users and Memos
- Full CRUD operations
- Swagger documentation
- Database migration support using `golang-migrate`
- Docker-based development environment
- Makefile for common tasks

---

## 🧑‍💻 How to Add a New CRUD Resource

To add a new resource (e.g., `User`, `Memo`), follow these steps:

### 1. Create Model

Create a struct in `domain/model/<resource>.go`:

```go
// domain/model/user.go
type User struct {
    ID        uint64    `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### 2. Add Migration File

```sh
make new-migration NAME=create_users
# Edit the generated SQL file in assets/migrations
```

### 3. Create Repository

File: `infrastructure/repository/<resource>_repository.go`

```go
func (r *resourceRepository) Create(db *gorm.DB, entity *model.Resource) error { return db.Create(entity).Error }
func (r *resourceRepository) List(db *gorm.DB) ([]model.Resource, error) { ... }
func (r *resourceRepository) Get(db *gorm.DB, id uint64) (*model.Resource, error) { ... }
func (r *resourceRepository) Update(db *gorm.DB, entity *model.Resource) error { ... }
func (r *resourceRepository) Delete(db *gorm.DB, id uint64) error { ... }
```

### 4. Create Usecase

File: `usecase/<resource>_usecase.go`

```go
func (u *resourceUsecase) Create(entity *model.Resource) error { return u.repo.Create(u.db, entity) }
```

### 5. Create Handler

File: `interface/handler/<resource>_handler.go`

```go
func (h *ResourceHandler) Create(c echo.Context) error {
    var entity model.Resource
    if err := c.Bind(&entity); err != nil { return util.ErrorJSON(...) }
    if err := h.usecase.Create(&entity); err != nil { ... }
    return c.JSON(http.StatusCreated, entity)
}
```

### 6. Register Route

File: `router/route.go`

```go
e.POST("/resources", handler.Create)
e.GET("/resources", handler.List)
e.GET("/resources/:id", handler.Get)
e.PUT("/resources/:id", handler.Update)
e.DELETE("/resources/:id", handler.Delete)
```

### 7. Add to DI (Dependency Injection)

Update the following files:

- `di/repositories.go`
- `di/usecases.go`
- `di/handlers.go`

---

## ⚙️ Makefile Command Reference

| Command                                | Description                                     |
| -------------------------------------- | ----------------------------------------------- |
| `make help`                            | Show help for all make commands                 |
| `make init-env`                        | Create `.env` from `.env.example` if not exists |
| `make check-env`                       | Validate required environment variables         |
| `make run`                             | Run the CLI app                                 |
| `make serve`                           | Run the HTTP server (`go run main.go serve`)    |
| `make build`                           | Build the Go binary                             |
| `make clean`                           | Delete the Go binary                            |
| `make docker-up`                       | Start Docker containers                         |
| `make docker-down`                     | Stop Docker containers                          |
| `make docker-volume-clean`             | Remove DB volume                                |
| `make docker-rebuild`                  | Rebuild containers without cache                |
| `make docker-exec`                     | Exec docker containers                          |
| `make fast-run`                        | Rebuild & reset DB & start server with logs     |
| `make launch`                          | Just start the server and show logs             |
| `make reset`                           | Stop, clean, and restart DB and app             |
| `make logs`                            | Tail app logs (`docker logs -f`)                |
| `make migrate-up`                      | Apply all migrations                            |
| `make migrate-down`                    | Rollback last migration                         |
| `make migrate-version`                 | Show current migration version                  |
| `make new-migration NAME=create_users` | Create new migration files                      |
| `make swag-init`                       | Generate Swagger docs                           |
| `make swag-open`                       | Open Swagger UI in browser                      |
| `make lint`                            | Run Go linters                                  |
| `make test`                            | Run tests with coverage                         |
| `make test-migrate`                    | Run migration-related tests                     |

---

## 🚀 Start Development

```sh
make init-env
make fast-run
open http://localhost:8080/swagger/index.html
```
