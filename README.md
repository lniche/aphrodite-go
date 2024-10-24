# Aphrodite Go API 快速开发脚手架

Aphrodite 是一个基于 [nunu](https://github.com/go-nunu/nunu) 开发的模板项目，旨在为 能够快速上手开发，理解框架使用流程，提供部署模板。

## 特性

| 技术                                | 说明                                         |
|-------------------------------------|--------------------------------------------|
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)       | SQL 模拟库，用于测试数据库交互                      |
| [lancet](https://github.com/duke-git/lancet/v2)             | 提供常用工具函数的库，简化 Go 开发                  |
| [gin](https://github.com/gin-gonic/gin)                     | 高性能的 Web 框架，基于 HTTP 的 RESTful API       |
| [sqlite](https://github.com/glebarez/sqlite)               | SQLite 数据库驱动，支持 GORM                     |
| [gocron](https://github.com/go-co-op/gocron)               | 简易的定时任务调度库                             |
| [jwt](https://github.com/golang-jwt/jwt/v5)                 | JWT 认证库，用于生成和验证 JSON Web Tokens      |
| [mock](https://github.com/golang/mock)                      | Go 的模拟对象生成工具                            |
| [wire](https://github.com/google/wire)                      | 依赖注入工具，用于自动化构建 Go 依赖关系           |
| [go-redis](https://github.com/redis/go-redis/v9)           | Redis 客户端库，支持 Redis 的异步操作             |
| [sonyflake](https://github.com/sony/sonyflake)              | 雪花算法实现的 ID 生成器                          |
| [viper](https://github.com/spf13/viper)                     | 配置管理库，支持多种格式的配置文件                |
| [testify](https://github.com/stretchr/testify)               | 测试框架，提供断言、模拟等功能                     |
| [files](https://github.com/swaggo/files)                    | Swagger 文件支持                                |
| [gin-swagger](https://github.com/swaggo/gin-swagger)       | 基于 Gin 的 Swagger UI 集成                       |
| [swag](https://github.com/swaggo/swag)                      | 自动生成 Swagger 文档的工具                      |
| [zap](https://github.com/uber-go/zap)                       | 高性能的日志库，支持结构化日志                    |
| [crypto](https://golang.org/x/crypto)                       | Go 的加密库，提供多种加密算法                     |
| [grpc](https://google.golang.org/grpc)                      | gRPC 框架，用于高效的远程调用                      |
| [lumberjack](https://gopkg.in/natefinch/lumberjack.v2)     | 日志轮转库，支持日志文件的管理                    |
| [gorm-mysql](https://gorm.io/driver/mysql)                  | GORM 的 MySQL 驱动                             |
| [gorm-postgres](https://gorm.io/driver/postgres)            | GORM 的 PostgreSQL 驱动                         |
| [gorm](https://gorm.io/gorm)                                 | 结构化数据的 ORM 框架，支持多种数据库              |

| 功能                  | 说明                                         |
|---------------------|--------------------------------------------|
| 登录授权功能        | 提供基础的用户登录和权限授权功能                    |
| 分布式锁            | 基于 Redis 实现的分布式锁                          |
| 中间件              | 包含认证、请求日志、跨域中间件                    |
| 实用封装            | 提供 AES、Hash、时间格式化等常用工具封装             |
| 统一输出方式        | 简单易用的 API Result 统一输出方式                  |

## 模块说明

- app => 应用模块
- pkg => 公共模块

## 本地运行

首先，确保你已经安装了 Go 语言环境。然后，可以通过以下步骤安装 Aphrodite：

```bash
git clone https://github.com/lniche/aphrodite-go.git
cd aphrodite
go mod tidy
