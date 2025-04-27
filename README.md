# Go + PostgreSQL Stocks API

A simple RESTful API built with Go and PostgreSQL to perform basic CRUD operations on a `stocks` database.

---

## Features

- Create a new stock
- Fetch all stocks
- Fetch a single stock by ID
- Update a stock by ID
- Delete a stock by ID
- Environment variable support for database configuration
- Clean project structure with Router and Middleware separation

---

## Technologies Used

- **Go** (Golang)
- **PostgreSQL**
- **Gorilla Mux** for routing
- **github.com/joho/godotenv** for environment variable loading

---

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name

2. Install dependencies

go mod init your-module-name
go mod tidy


3. Set up PostgreSQL
Install PostgreSQL if you haven't already.

Start PostgreSQL server.

Create a database:

createdb myappdb


4. Create the stocks table

CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    company VARCHAR(100) NOT NULL
);


5. Create .env file
In the project root, create a .env file:

POSTGRES_URL="postgres://your-username:your-password@localhost:5432/your-database-name?sslmode=disable"


Running the Server
bash
Copy
Edit
go run main.go
Server will start on:

http://localhost:8080
