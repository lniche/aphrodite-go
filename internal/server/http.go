package server

import (
	apiV1 "aphrodite-go/api/v1"
	"aphrodite-go/docs"
	"aphrodite-go/internal/handler"
	"aphrodite-go/internal/middleware"
	"aphrodite-go/pkg/jwt"
	"aphrodite-go/pkg/log"
	"aphrodite-go/pkg/server/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	redis *redis.Client,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))
	s.GET("/swagger-ui/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.RequestIDMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		// middleware.SignMiddleware(logger),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.Ok(ctx, "Thank you for using Aphrodite!")
	})
	s.GET("/ping", func(ctx *gin.Context) {
		apiV1.Ok(ctx, "pong")
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/login", authHandler.Login)
			noAuthRouter.POST("/logout", authHandler.Logout)
			noAuthRouter.POST("/send-code", authHandler.SendVerifyCode)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger, redis))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetUser)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger, redis))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateUser)
			strictAuthRouter.DELETE("/user", userHandler.DeleteUser)
		}
	}

	return s
}
