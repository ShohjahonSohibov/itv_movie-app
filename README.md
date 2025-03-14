# Movie Management API

A RESTful API for managing movies with PostgreSQL database.

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15
- Docker (optional)

## Running the Project
- make swag
- make run

### Local Development
- GET /api/v1/movies - List all movies
- GET /api/v1/movies/{id} - Get movie by ID
- POST /api/v1/movies - Create new movie
- PUT /api/v1/movies/{id} - Update movie
- DELETE /api/v1/movies/{id} - Delete movie

1. Start PostgreSQL:
```bash
docker compose up postgres -d
```

# http://localhost:8080/swagger/index.html