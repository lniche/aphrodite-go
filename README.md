# Aphrodite Go API 快速开发脚手架

Aphrodite 是一个基于 [nunu](https://github.com/go-nunu/nunu) 开发的模板项目，旨在帮助开发者快速上手，深入理解框架的使用流程。该项目提供了全面的示例代码和配置，涵盖了常见的开发场景，以便于学习和实践。此外，Aphrodite 还包含容器部署模板，使得项目在现代云环境中能够轻松部署与管理，助力开发者高效构建和发布应用。

## 技术栈

| 技术                                                   | 说明                                        |
| ------------------------------------------------------ | ------------------------------------------- |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)   | SQL 模拟库，用于测试数据库交互              |
| [lancet](https://github.com/duke-git/lancet/v2)        | 提供常用工具函数的库，简化 Go 开发          |
| [gin](https://github.com/gin-gonic/gin)                | 高性能的 Web 框架，基于 HTTP 的 RESTful API |
| [sqlite](https://github.com/glebarez/sqlite)           | SQLite 数据库驱动，支持 GORM                |
| [gocron](https://github.com/go-co-op/gocron)           | 简易的定时任务调度库                        |
| [jwt](https://github.com/golang-jwt/jwt/v5)            | JWT 认证库，用于生成和验证 JSON Web Tokens  |
| [mock](https://github.com/golang/mock)                 | Go 的模拟对象生成工具                       |
| [wire](https://github.com/google/wire)                 | 依赖注入工具，用于自动化构建 Go 依赖关系    |
| [go-redis](https://github.com/redis/go-redis/v9)       | Redis 客户端库，支持 Redis 的异步操作       |
| [sonyflake](https://github.com/sony/sonyflake)         | 雪花算法实现的 ID 生成器                    |
| [viper](https://github.com/spf13/viper)                | 配置管理库，支持多种格式的配置文件          |
| [testify](https://github.com/stretchr/testify)         | 测试框架，提供断言、模拟等功能              |
| [files](https://github.com/swaggo/files)               | Swagger 文件支持                            |
| [gin-swagger](https://github.com/swaggo/gin-swagger)   | 基于 Gin 的 Swagger UI 集成                 |
| [swag](https://github.com/swaggo/swag)                 | 自动生成 Swagger 文档的工具                 |
| [zap](https://github.com/uber-go/zap)                  | 高性能的日志库，支持结构化日志              |
| [crypto](https://golang.org/x/crypto)                  | Go 的加密库，提供多种加密算法               |
| [grpc](https://google.golang.org/grpc)                 | gRPC 框架，用于高效的远程调用               |
| [lumberjack](https://gopkg.in/natefinch/lumberjack.v2) | 日志轮转库，支持日志文件的管理              |
| [gorm-mysql](https://gorm.io/driver/mysql)             | GORM 的 MySQL 驱动                          |
| [gorm-postgres](https://gorm.io/driver/postgres)       | GORM 的 PostgreSQL 驱动                     |
| [gorm](https://gorm.io/gorm)                           | 结构化数据的 ORM 框架，支持多种数据库       |

## 特性

- **用户认证与授权**：提供基础的用户登录和权限授权功能。
- **分布式锁**：基于 Redis 实现的分布式锁，保证分布式环境下的资源安全。
- **中间件支持**：内置常用的中间件，包括认证、请求日志、跨域处理等。
- **统一输出格式**：提供简单易用的 API Result 统一输出方式，标准化 API 响应格式，提升接口一致性。
- **API 模块化设计**：支持模块化的 API 设计，易于扩展和维护。
- **Swagger 文档集成**：自动生成 API 文档，便于前端开发和测试。

## 目录结构

```
.
├── api/                  # 出入参数
├── cmd/                  # 应用程序的入口，包含不同的子命令
├── config/               # 配置文件，存储应用的各种配置项
├── docs/               # 文档，swagger文件
├── deploy/               # 部署相关的文件，如 Dockerfile、docker-compose.yml 等
├── internal/             # 应用程序的主要代码，按分层架构组织
├── pkg/                  # 公共模块，包含共享的代码，如配置、日志、HTTP 等功能
├── scripts/               # 脚本文件，用于部署和其他自动化任务
├── storage/              # 存储文件，如日志文件和数据库文件
└── main.go               # 入口文件
```

## 本地运行

```bash
# 1. 克隆项目代码库
git clone https://github.com/lniche/aphrodite-go.git
cd aphrodite-go

# 2. 配置文件
cd config
mv config.yml.example config.yml

# 3. 处理依赖
# 确保你已经安装了 Go
go mod tidy

# 4. 初始化数据库
deploy/db.sql

# 5. 启动服务
# 确保你的环境中安装了 nunu
go install github.com/go-nunu/nunu@latest
nunu run
```

## 贡献

如果你有任何建议或想法，欢迎创建 Issue 或直接提交 Pull Request。

1. Fork 这个仓库。
2. 创建一个新的分支：git checkout -b feature/your-feature
3. 提交你的更改：git commit -m 'Add new feature'
4. 推送到你的分支：git push origin feature/your-feature
5. 提交 Pull Request。

## 许可证

该项目遵循 MIT 许可证。

## 鸣谢

特别感谢所有贡献者和支持者，您的帮助对我们至关重要！
