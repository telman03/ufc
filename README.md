# UFC Fight Tracking API

## Overview

The UFC Fight Tracking API is a RESTful API designed to manage and track UFC fighters, their rankings, and user favorites. This API allows users to register, log in, and manage their favorite fighters while providing detailed information about fighters and their rankings.

## Features

- User registration and authentication
- Manage favorite fighters
- Search and filter fighters based on various criteria
- Retrieve fighter rankings by weight class
- Swagger documentation for easy API exploration

## Technologies Used

- Go (Golang)
- Echo framework for building the API
- GORM for ORM (Object Relational Mapping)
- PostgreSQL for the database
- JWT (JSON Web Tokens) for authentication
- Swagger for API documentation

## Getting Started

### Prerequisites

- Go (1.16 or later)
- PostgreSQL
- Go modules

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/telman03/ufc.git
   cd ufc
   ```

2. Create a `.env` file in the root directory and add your database configuration:

   ```env
   DB_HOST=localhost
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_PORT=5432
   JWT_SECRET=your_jwt_secret
   ```

3. Install the required Go packages:

   ```bash
   go mod tidy
   ```

4. Run the database migrations:

   ```bash
   go run src/backend/main.go
   ```

### Running the API

To start the API server, run:

```bash
go run src/backend/main.go
```

The server will start on `http://localhost:8080`.

### API Endpoints

#### User Authentication

- **Register User**
  - `POST /register`
  - Request Body: `{"username": "testuser", "email": "user@example.com", "password": "123456"}`
  
- **Login User**
  - `POST /login`
  - Request Body: `{"email": "user@example.com", "password": "123456"}`

#### Fighter Management

- **Search Fighters**
  - `GET /fighters`
  - Query Parameters: `name`, `stance`, `weight`, `wins`, `losses`, `limit`, `offset`

- **Manage Favorites**
  - **Add Favorite**: `POST /favorites`
    - Request Body: `{"fighter_id": 1}`
  - **Remove Favorite**: `DELETE /favorites/{fighter_id}`
  - **List Favorites**: `GET /favorites`

#### Rankings

- **Get Rankings by Weight Class**
  - `GET /rankings`
  - Query Parameter: `weightclass`

#### Swagger Documentation

The API is documented using Swagger. You can access the documentation at:

```
http://localhost:8080/swagger/index.html
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Echo](https://echo.labstack.com/) - High performance, minimalist Go web framework
- [GORM](https://gorm.io/) - ORM for Golang
- [Swagger](https://swagger.io/) - API documentation tool
