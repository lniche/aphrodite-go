# Aphrodite Go API 快速开发脚手架

Aphrodite 是一个基于 [nunu](https://github.com/go-nunu/nunu) 开发的模板项目，旨在帮助开发者快速上手，深入理解框架的使用流程。该项目提供了全面的示例代码和配置，涵盖了常见的开发场景，以便于学习和实践。此外，Aphrodite 还包含容器部署模板，使得项目在现代云环境中能够轻松部署与管理，助力开发者高效构建和发布应用。

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
- cmd => 应用程序的入口，包含不同的子命令
- config => 配置文件，存储应用的各种配置项
- deploy => 部署相关的文件，如 Dockerfile、docker-compose.yml 等
- internal => 应用程序的主要代码，按照分层架构进行组织
  - handler (or controller) => 处理 HTTP 请求，调用业务逻辑层的服务，返回 HTTP 响应
  - server (or router) => HTTP 服务器，启动 HTTP 服务，监听端口，处理 HTTP 请求
  - service (or logic) => 服务，实现具体的业务逻辑，调用数据访问层 repository
  - model (or entity) => 数据模型，定义了业务逻辑层需要的数据结构
  - repository (or dao) => 数据访问对象，封装了数据库操作，提供了对数据的增删改查
- pkg => 公共模块，包含共享的代码，如配置、日志、HTTP 等功能
- script => 脚本文件，用于部署和其他自动化任务
- storage => 存储文件，如日志文件和数据库文件
- test => 测试代码，包含单元测试和集成测试

## 本地运行

```bash
# 1. 初始化数据库
# 例如：将数据库脚本放在 deploy/db.sql 路径下

# 2. 克隆项目代码库
git clone https://github.com/lniche/aphrodite-go.git
cd aphrodite-go

# 3. 处理依赖
# 使用 go mod tidy 来清理并下载项目所需的所有依赖
go mod tidy

# 4. 配置文件
# 将配置示例文件重命名为 config.yml
mv config.yml.example config.yml

# 5. 编辑配置文件（如有需要）
# 使用你喜欢的文本编辑器打开 config.yml
# 例如使用 vim：
vim config.yml
# 根据你的需求修改配置参数（如数据库连接、服务端口等）

# 6. 启动服务
# 确保你的环境中安装了 nunu
nunu run
```