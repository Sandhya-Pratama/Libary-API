****Library Management System**
**
A scalable library management system built using a microservices architecture. This backend-only project provides RESTful APIs for managing books, authors, categories, users, and book borrowing operations. The project is developed with Golang, PostgreSQL for data storage, gRPC for inter-service communication, and Docker for containerization.

**Features**

- Book Management: CRUD operations for managing books.
- Author Management: CRUD operations for managing authors.
- Category Management: CRUD operations for managing book categories.
- User Management: CRUD operations for managing library users.
- Borrowing and Returning: Borrowing and returning books with stock management.
- JWT Authentication: Secure endpoints with JSON Web Token (JWT) for authentication and authorization.

**Technology Stack**
-Programming Language: Go (Golang)
-Database: PostgreSQL (each service has its own database)
-Architecture: Microservices
-Inter-service Communication: gRPC
-Authentication: JWT
-Containerization: Docker
-Orchestration: Docker Compose

**Microservices**
The application is divided into several microservices, each with its own database:

-User Service: Manages users and authentication.
-Author Service: Manages authors' information.
-Category Service: Manages book categories.
-Book Service: Manages book data and stock.
-Borrowing Service: Manages borrowing and returning of books.

**Directory Structure**
library-management-system/
├── user-service/
│   ├── cmd/server/main.go
│   ├── common/
│   ├── db/migrations/
│   ├── entity/
│   ├── internal/
│   ├── Dockerfile
│   └── .env
├── author-service/
│   └── (similar structure as user-service)
├── category-service/
│   └── (similar structure as user-service)
├── book-service/
│   └── (similar structure as user-service)
├── borrowing-service/
│   └── (similar structure as user-service)
├── docker-compose.yml
├── README.md
└── .env

**Prerequisites**
Golang: v1.18 or higher
Docker and Docker Compose
PostgreSQL (used in Docker Compose for local development)

**Getting Started**
1. Clone the Repository
Start to clone repository in terminal using: 
git clone https://github.com/Sandhya-Pratama/Libary-API.git
cd Libary-API

2. Set Up Environment Variables
Each service has its own .env file for configuration. Below is an example of environment variables needed for each service.

Example .env for user-service
DB_HOST=user-db
DB_PORT=5432
DB_USER=user
DB_PASSWORD=user_password
DB_NAME=user_db
JWT_SECRET=your_jwt_secret_key

3. Start Services Using Docker Compose
This project uses Docker Compose for orchestration. To start all services, run the following command:
docker-compose up --build

4. Accessing the Services
Once the services are running, you can access the APIs directly via the following ports:

-User Service: http://localhost:8080/api/v1/users
-Author Service: http://localhost:8081/api/v1/authors
-Category Service: http://localhost:8082/api/v1/categories
-Book Service: http://localhost:8083/api/v1/books
-Borrowing Service: http://localhost:8084/api/v1/borrowings

5. gRPC Communication
Each service communicates with others via gRPC for certain functionalities. Ensure gRPC communication settings in .env files are correctly configured with the appropriate host and port.

**API Documentation**
Below are some sample endpoints for each service.

User Service
-POST /api/v1/users/register: Register a new user.
-POST /api/v1/users/login: User login and receive JWT.

Book Service
-POST /api/v1/books: Add a new book.
-GET /api/v1/books/{id}: Retrieve book details.
-PUT /api/v1/books/{id}: Update book information.
-DELETE /api/v1/books/{id}: Delete a book.

Borrowing Service
-POST /api/v1/borrowings: Borrow a book.
-POST /api/v1/borrowings/return: Return a book.
