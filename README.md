# GO-y-DO

GO-y-DO is a simple TODO manager built using GOLang, Gin, GORM, and PostgreSQL. It supports creating, retrieving, updating, and deleting tasks. The project also includes API documentation generated using Swagger.

## Getting Started

### Quick Start

```bash
docker-compose up --build -d
```

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18+)
- [PostgreSQL](https://www.postgresql.org/download/)
- [GORM](https://gorm.io/)
- [Gin](https://github.com/gin-gonic/gin)
- [Swagger](https://swagger.io/tools/swagger-ui/)

### Installation

1. **Set up environment variables:**

   Create a `.env` file in the root directory of your project by copying the provided `env.example` file:

   ```sh
   cp env.example .env
   ```

   Edit the `.env` file to match your PostgreSQL configuration.

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Generate Swagger documentation:**

   Run the following command to generate Swagger documentation in the `src/docs` directory:

   ```sh
   swag init -o src/docs
   ```

### Running the Project

1. **Run the server:**

   ```sh
   go run main.go
   ```

   The server will start on `http://localhost:3000`.

2. **Access Swagger UI:**

   Visit `http://localhost:3000/swagger/index.html` to view and interact with the API documentation.

### Project Configuration

Database configuration is managed via environment variables. The `.env` file should look like this:

```
DATABASE_URL=postgres://<username>:<password>@<host>:<port>/<dbname>
```
