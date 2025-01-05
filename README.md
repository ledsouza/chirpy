# Chirpy - A Twitter-like API built with Golang

Chirpy is a REST API project built with Golang, inspired by Twitter, offering core functionalities like user creation, authentication, chirping (posting), and basic administrative features.  It serves as a practical example of building a backend API, demonstrating the use of various technologies and best practices.

## Why Chirpy?

This project is valuable for developers looking to:

* **Learn Golang API Development:** Explore building a RESTful API with Go, handling routing, database interactions, and authentication.
* **Understand Database Migrations:**  The project utilizes `goose` for database migrations, showcasing how to manage database schema changes effectively.
* **Implement JWT Authentication:**  Chirpy incorporates JWT (JSON Web Token) for secure user authentication and authorization.
* **Practice with PostgreSQL and SQLX:** See how to interact with a PostgreSQL database using `psql` for code generation from SQL queries.
* **Study API Metrics and Monitoring:**  Basic administrative endpoints provide insights into API usage and allow for resetting data in a development environment.

## Features

* **User Management:** Create, update, and log in users.
* **Chirping:** Post and retrieve chirps (tweets), including listing chirps by author and sorting.
* **Authentication:** Secure API access with JWT, including refresh and revoke token functionality.
* **Administrative Features:**  Metrics tracking and data reset capabilities (dev environment only).
* **Database Migrations:** Managed using `goose`.
* **PostgreSQL Integration:**  Leveraging `psql` for SQL query-based code generation.


## Installation and Running

### Prerequisites

* **Go:** Make sure you have Go installed on your system.  You can download it from [https://go.dev/dl/](https://go.dev/dl/).
* **PostgreSQL:**  Install and set up a PostgreSQL database.
* **Goose:** Install the `goose` CLI tool for database migrations. ([https://github.com/pressly/goose](https://github.com/pressly/goose))
* **psql:** Ensure you have the `psql` command-line utility available for interacting with your PostgreSQL database.


### Steps

1. **Clone the Repository:**

```bash
git clone https://github.com/ledsouza/chirpy.git
cd chirpy
```

2. **Set Up Environment Variables:**

Create a `.env` file in the root directory and add the following environment variables:

```
DB_URL=your_postgresql_connection_string  # Example: postgres://user:password@host:port/database?sslmode=disable
PLATFORM=dev  # Or "prod"
JWT_SECRET_KEY=your_secret_key
POLKA_KEY=your_polka_api_key
```

3. **Run Database Migrations:**

```bash
goose postgres "postgres://username:password@localhost:5432/chirpy" up 
```

4. **Run the Application:**

```bash
go run main.go
```

The API will start running on port 8080.


## Contributing

Contributions are welcome!  Please feel free to open issues and submit pull requests.
