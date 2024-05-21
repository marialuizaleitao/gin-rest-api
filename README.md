# Beatles Fan Club API

The Beatles Fan Club API is a simple RESTful API built using Go (Golang) and Gin framework. It allows users to interact with a database of Beatles fans, perform CRUD operations on member data, and receive personalized hello messages.

## Routes

The API provides the following routes:

- `GET /members`: Retrieves all Beatles fans.
- `GET /:name`: Displays a personalized hello message.
- `POST /members`: Adds a new member to the fan club.
- `GET /members/:id`: Retrieves a member by their ID.
- `DELETE /members/:id`: Deletes a member from the fan club.
- `PATCH /members/:id`: Updates a member's information.

## Getting Started

To run the API locally, follow these steps:

1. Clone this repository:

    ```bash
    git clone https://github.com/marialuizaleitao/gin-rest-api.git
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Start the server:

    ```bash
    go run main.go
    ```

   The API will be available at [http://localhost:8080](http://localhost:8080).

## Technologies Used

The Beatles Fan Club API uses the following technologies:

- **Go (Golang)**: Programming language used to build the API.
- **Gin**: Web framework for Go, used for handling HTTP requests and responses.
- **GORM**: ORM library for Go, used for interacting with the database.
- **Swagger**: Documentation tool for RESTful APIs, used for generating API documentation.

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.
