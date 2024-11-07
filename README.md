# Aphrodite Gin API Scaffold

Aphrodite is a template project based on [Nunu](https://github.com/go-nunu/nunu) to help developers get started quickly and gain a deep understanding of the framework's usage process. The project provides comprehensive sample code and configuration, covering common development scenarios for easy learning and practice. In addition, Aphrodite also includes container deployment templates, making the project easy to deploy and manage in a modern cloud environment, helping developers to efficiently build and release applications.

## Tech Stack

| Technology                                             | Description                                                                          |
| ------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)   | SQL simulation library for testing database interactions                             |
| [lancet](https://github.com/duke-git/lancet/v2)        | A library that provides common tool functions to simplify Go development             |
| [gin](https://github.com/gin-gonic/gin)                | High-performance web framework, HTTP-based RESTful API                               |
| [sqlite](https://github.com/glebarez/sqlite)           | SQLite database driver, supports GORM                                                |
| [gocron](https://github.com/go-co-op/gocron)           | Simple scheduled task scheduling library                                             |
| [jwt](https://github.com/golang-jwt/jwt/v5)            | JWT authentication library for generating and verifying JSON Web Tokens              |
| [mock](https://github.com/golang/mock)                 | Mock object generation tool for Go                                                   |
| [wire](https://github.com/google/wire)                 | Dependency injection tool for automatically building Go dependencies                 |
| [go-redis](https://github.com/redis/go-redis/v9)       | Redis client library, supporting asynchronous operations of Redis                    |
| [sonyflake](https://github.com/sony/sonyflake)         | ID generator implemented by the snowflake algorithm                                  |
| [viper](https://github.com/spf13/viper)                | Configuration management library, supporting configuration files in multiple formats |
| [testify](https://github.com/stretchr/testify)         | Test framework, providing assertion, simulation and other functions                  |
| [files](https://github.com/swaggo/files)               | Swagger file support                                                                 |
| [gin-swagger](https://github.com/swaggo/gin-swagger)   | Swagger UI integration based on Gin                                                  |
| [swag](https://github.com/swaggo/swag)                 | Tool for automatically generating Swagger documents                                  |
| [zap](https://github.com/uber-go/zap)                  | High-performance logging library, supporting structured logging                      |
| [crypto](https://golang.org/x/crypto)                  | Go's encryption library, providing multiple encryption algorithms                    |
| [grpc](https://google.golang.org/grpc)                 | gRPC framework for efficient remote calls                                            |
| [lumberjack](https://gopkg.in/natefinch/lumberjack.v2) | Log rotation library, supporting log file management                                 |
| [gorm-mysql](https://gorm.io/driver/mysql)             | GORM's MySQL driver                                                                  |
| [gorm-postgres](https://gorm.io/driver/postgres)       | PostgreSQL driver for GORM                                                           |
| [gorm](https://gorm.io/gorm)                           | ORM framework for structured data, supporting multiple databases                     |

## Features

- **User authentication and authorization**: Provides basic user login and permission authorization functions.
- **Distributed lock**: Distributed lock based on Redis to ensure resource security in a distributed environment.
- **Middleware support**: Built-in commonly used middleware, including authentication, request log, cross-domain processing, etc.
- **Unified output format**: Provides a simple and easy-to-use API Result unified output method, standardizes the API response format, and improves interface consistency.
- **API modular design**: Supports modular API design, which is easy to expand and maintain.
- **Swagger document integration**: Automatically generates API documents for front-end development and testing.

## Directory structure

```
.
├── api/ # Input and output parameter definitions, including all request and response structures
├── bin/ # Executable scripts
├── cmd/ # Application entry, including different subcommands, such as startup, migration, etc.
├── config/ # Configuration file, storing application configuration items (such as database, third-party services, etc.)
├── database/ # Database related
├── deploy/ # Deployment related files
├── docs/ # Project documentation
├── internal/ # Application core code
├── pkg/ # Public module
├── storage/ # File storage
├── tests/ # Test files
└── README.md # Project description
```

## Local operation

```bash
# 1. Clone the project code base
git clone https://github.com/lniche/aphrodite-go.git
cd aphrodite-go

# 2. Configuration file
cd config
mv config.yml.example config.yml

# 3. Handle dependencies
# Make sure you have installed the Go environment
go mod tidy

# 4. Initialize the database
database/db.sql

# 5. Start the service
# Make sure nunu is installed in your environment
go install github.com/go-nunu/nunu@latest
nunu run
```

## Repo Activity

![Alt](https://repobeats.axiom.co/api/embed/75f6227f2c9b38043ecc5b2c0c5dfacd5cd373cb.svg "Repobeats analytics image")


## Contribution

If you have any suggestions or ideas, please create an Issue or submit a Pull Request directly.

1. Fork this repository.
2. Create a new branch: git checkout -b feature/your-feature
3. Commit your changes: git commit -m 'Add new feature'
4. Push to your branch: git push origin feature/your-feature
5. Submit a Pull Request.

## License

This project is licensed under the MIT License.

## Acknowledgements

Special thanks to all contributors and supporters, your help is essential to us!
