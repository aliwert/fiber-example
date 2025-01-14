# Fiber Service

- This repository contains a Go Fiber application that demonstrates a simple setup with database connectivity and various handlers for managing users, cars, customers, and more.

## Project Structure

- `cmd/main.go`: Entry point of the application. It sets up the Fiber app, loads environment variables, and starts the server.
- `database/database.go`: Handles the database connection and initializes the required tables.
- `handlers/`: Contains handlers for different routes like authentication and car management.
- `middleware/auth.go`: Contains middleware for protecting routes using JWT authentication.
- `models/`: Contains the data models used in the application.
- `routes/routes.go`: Defines the API routes for the application, including authentication, cars, customers, maintenance, and rentals.

## Setup and Installation

1. Clone the Repository

```
git clone https://github.com/aliwert/fiber-example.git

cd fiber-example
```

2. Install Dependencies

Make sure you have [Go](https://go.dev/dl/) installed. Then, run:

```sh
go mod tidy
```

3. Environment Variables
   Create a .`env` file in the root directory and add the following environment variables:

```env
DB_HOST=<your_db_host>
DB_PORT=<your_db_port>
DB_USER=<your_db_user>
DB_PASSWORD=<your_db_password>
DB_NAME=<your_db_name>
JWT_SECRET=<your_jwt_secret>
PORT=<server_port>
```

4. Run the Application

```sh
go run cmd/main.go
```

## Usage

### Register a User

Send a POST request to /register with the following JSON body:

```json
{
  "email": "user@example.com",
  "password": "your_password"
}
```

### Login

Send a POST request to /login with the following JSON body:

```json
{
  "email": "user@example.com",
  "password": "your_password"
}
```
