# Go Authentication with JWT and Refresh Tokens

## Overview

This is a Go application that implements authentication using JWT (JSON Web Tokens) with support for refresh tokens. The app follows Clean Architecture principles to separate concerns between the domain, use cases, repository, and delivery layers. It interacts with a PostgreSQL database for storing user information and token data.

The app also provides a set of HTTP endpoints for user authentication, including login, logout, token refresh, and user registration.

The app is Dockerized for easy deployment and scaling.

## Technologies Used

- **Go**: The programming language used for backend development.
- **PostgreSQL**: Relational database to store user and token information.
- **JWT**: JSON Web Tokens used for securing API endpoints.
- **Docker**: For containerizing the application.
- **Clean Architecture**: To ensure a scalable, maintainable, and testable codebase.

## Folder Structure

```plaintext
.
├── cmd
│   ├── main.go            # Application entry point
├── config
│   ├── config.go          # Configuration and environment variables
├── internal
│   ├── domain
│   │   ├── models.go      # Data models for users, tokens, etc.
│   │   ├── jwt_utils.go   # JWT generation and validation logic
│   ├── repository
│   │   ├── user_repo.go   # Database interaction for users and tokens
│   ├── usecase
│   │   ├── auth_usecase.go # Business logic for authentication
│   ├── delivery
│       ├── http
│           ├── auth_handler.go # HTTP handlers for authentication routes
├── migrations
│   ├── <migration-files>  # PostgreSQL migration files
├── go.mod
├── go.sum
└── README.md
```

## Endpoints

1. **POST /auth/login**
   - Description: Logs in a user and returns a JWT and refresh token.
   - Request Body: `{"username": "user", "password": "password"}`
   - Response:
     ```json
     {
       "access_token": "jwt_token_here",
       "refresh_token": "refresh_token_here"
     }
     ```
   
2. **POST /auth/register**
   - Description: Registers a new user.
   - Request Body: `{"username": "new_user", "password": "password"}`
   - Response: 
     ```json
     {
       "message": "User registered successfully"
     }
     ```

3. **POST /auth/refresh**
   - Description: Refreshes the access token using a valid refresh token.
   - Request Body: `{"refresh_token": "valid_refresh_token_here"}`
   - Response:
     ```json
     {
       "access_token": "new_jwt_token_here"
     }
     ```

4. **POST /auth/logout**
   - Description: Logs out a user by invalidating the refresh token.
   - Request Body: `{"refresh_token": "valid_refresh_token_here"}`
   - Response: 
     ```json
     {
       "message": "Logged out successfully"
     }
     ```

## Technologies and Libraries

- **Go 1.18+**
- **PostgreSQL**: Used for storing user data and tokens.
- **JWT-Go**: Library to create and verify JWT tokens.
- **Gorilla Mux**: HTTP router for Go.
- **Docker**: For containerizing the app.

## Dockerization

To run the application using Docker, follow these steps:

1. **Build the Docker image**:

   ```bash
   docker build -t go-auth-app .
   ```

2. **Run the container**:

   ```bash
   docker run -d -p 8080:8080 --env-file .env go-auth-app
   ```

   This will start the application inside a Docker container, listening on port `8080`.

3. **Stop the container**:

   ```bash
   docker stop <container_id>
   ```

4. **Docker Compose**:

   If you also want to run PostgreSQL as a service using Docker Compose, you can create a `docker-compose.yml` file like this:

   ```yaml
   version: '3.8'
   services:
     app:
       build: .
       ports:
         - "8080:8080"
       env_file:
         - .env
       depends_on:
         - db
     db:
       image: postgres:13
       environment:
         POSTGRES_USER: user
         POSTGRES_PASSWORD: password
         POSTGRES_DB: authdb
       ports:
         - "5432:5432"
   ```

   Then, run:

   ```bash
   docker-compose up --build
   ```

## Setup and Configuration

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/go-auth-app.git
   cd go-auth-app
   ```

2. Set up the PostgreSQL database by running the migration scripts. This will create the necessary tables for users and tokens in the database.

3. Create a `.env` file with the following configuration:

   ```plaintext
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=user
   DB_PASSWORD=password
   DB_NAME=authdb
   JWT_SECRET=your_secret_key
   JWT_EXPIRATION=15m
   JWT_REFRESH_EXPIRATION=7d
   ```

4. Run the application:

   ```bash
   go run cmd/main.go
   ```

## Database Migrations

Migrations are used to initialize and update the database schema. You can add migration files under the `migrations` folder.

You can use `golang-migrate` or any other migration tool of your choice.

