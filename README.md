
# Backend Development Project - REST API with PostgreSQL, MySQL, Redis, Gin, and Gorm

This project is a REST API developed as part of the Backend Development course at Iran University of Science and Technology (IUST), CESA College. The API uses **PostgreSQL** and **MySQL** as databases, **Redis** for caching, and is built using the **Gin framework** and **Gorm** for ORM. The entire application is Dockerized for easy deployment and scalability.

## Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup and Installation](#setup-and-installation)
- [API Endpoints](#api-endpoints)
- [Dockerization](#dockerization)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

---

## Features
- **Dual Database Support**: The API supports both PostgreSQL and MySQL databases, allowing flexibility in database choice.
- **Caching with Redis**: Implements Redis for caching frequently accessed data to improve performance.
- **RESTful API**: Follows REST principles for clean and scalable API design.
- **Dockerized**: The entire application is containerized using Docker for easy deployment and testing.
- **Gin Framework**: Utilizes the high-performance Gin framework for building the API.
- **Gorm ORM**: Uses Gorm for database operations, providing a clean and efficient way to interact with databases.

---

## Technologies Used
- **Backend Framework**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [Gorm](https://gorm.io/)
- **Databases**: PostgreSQL, MySQL
- **Caching**: Redis
- **Containerization**: Docker
- **Programming Language**: Go (Golang)
- **Other Tools**: Git, GitHub, Docker Compose

---

## Setup and Installation

### Prerequisites
- Docker and Docker Compose installed on your machine.
- Git for version control.

### Steps to Run the Project
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
   ```

2. **Set Up Environment Variables**:
   - Create a `.env` file in the root directory and add the necessary environment variables (e.g., database credentials, Redis URL, etc.).
   Example:
   ```env
   POSTGRES_USER=your_user
   POSTGRES_PASSWORD=your_password
   POSTGRES_DB=your_db
   REDIS_URL=redis://localhost:6379
   ```

3. **Build and Run with Docker**:
   ```bash
   docker-compose up --build
   ```
   This command will build the Docker images and start the containers for the API, databases, and Redis.

4. **Access the API**:
   - The API will be running at `http://localhost:8080`.

---

## API Endpoints
Here are some of the key endpoints available in the API:

- **GET /users**: Fetch all users.
- **POST /users**: Create a new user.
- **GET /users/{id}**: Fetch a specific user by ID.
- **PUT /users/{id}**: Update a user by ID.
- **DELETE /users/{id}**: Delete a user by ID.

For a complete list of endpoints, refer to the [API Documentation](#) (if applicable).

---

## Dockerization
The project is fully Dockerized, with separate containers for:
- The Go application (API)
- PostgreSQL database
- MySQL database
- Redis cache

The `docker-compose.yml` file defines the services and their configurations. You can easily modify it to suit your needs.

---

## Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Push your branch and open a pull request.

---

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgements
- **Instructor**: [Instructor Name] at Iran University of Science and Technology, CESA College.
- **Gin Framework**: [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **Gorm ORM**: [https://gorm.io/](https://gorm.io/)

---

This project demonstrates my skills in backend development, database management, caching, and containerization. It is a great addition to my resume and showcases my ability to build scalable and efficient REST APIs.

---

### How to Add This to Your Resume
You can include this project in the **Projects** section of your resume with the following description:

**Backend Development Project - REST API with PostgreSQL, MySQL, Redis, and Docker**  
- Developed a RESTful API using the Gin framework and Gorm ORM, supporting both PostgreSQL and MySQL databases.  
- Implemented Redis for caching to improve API performance.  
- Dockerized the application for easy deployment and scalability.  
- Technologies: Go (Golang), Gin, Gorm, PostgreSQL, MySQL, Redis, Docker.  
- [GitHub Repository](#) (link to your repo)

---
## Development Commands

This project includes a `Makefile` to simplify common tasks. Here are some useful commands:

- **Start the application**: `make all-up ENV_FILE=.env`
- **Stop the application**: `make all-down ENV_FILE=.env`
- **Run tests**: `make test TEST_API_DIR=api TEST_SERVICE_DIR=service`
- **Run database migrations**: `make migrate-up ENV_FILE=.env`
- **Clean up Docker resources**: `make clean ENV_FILE=.env`
---
## Configuration

This project uses environment variables for configuration. Create a `.env` file in the root directory and add the following variables:

```env
# Database Configuration
DB_DRIVER=postgres
DB_HOST=db
DB_PORT=5432
DB_USER=library
DB_PASSWORD=pass-example
DB_NAME=example

# Redis Configuration
CACHE_HOST=redis
CACHE_PORT=6379
CACHE_PASSWORD=pass-example
CACHE_DB=0

# HTTP Server Configuration
HTTP_URL=localhost
HTTP_PORT=8080

# Docker Configuration
DOCKER_HUB_USERNAME=example
IMAGE_TAG=alpinelinux/golang
HOST_PORT=8080
CONTAINER_PORT=8080

# Application Configuration
GIN_MODE=debug
APP_ENV=development
LOG_LEVEL=debug
MIGRATION_PATH=./migrations
This README and resume description will help you showcase your project effectively to potential employers or collaborators. Good luck! 🚀
