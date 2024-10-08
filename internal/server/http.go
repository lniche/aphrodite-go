package server

import (
	apiV1 "aphrodite-go/api/v1"
	"aphrodite-go/docs"
	"aphrodite-go/internal/handler"
	"aphrodite-go/internal/middleware"
	"aphrodite-go/pkg/jwt"
	"aphrodite-go/pkg/log"
	"aphrodite-go/pkg/server/http"
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
	userFeedbackHandler *handler.UserFeedbackHandler,
	userAddressHandler *handler.UserAddressHandler,
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
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using aphrodite!",
		})
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/auth/login", authHandler.Login)
			noAuthRouter.POST("/auth/send-code", authHandler.SendVerifyCode)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger, redis))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
			noStrictAuthRouter.GET("/user/address/:id", userAddressHandler.GetUserAddress)
			noStrictAuthRouter.GET("/user/address", userAddressHandler.GetUserAddresses)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger, redis))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
			strictAuthRouter.POST("/user/feedback", userFeedbackHandler.CreateUserFeedback)
			strictAuthRouter.POST("/user/address", userAddressHandler.CreateUserAddress)
			strictAuthRouter.PUT("/user/address", userAddressHandler.UpdateUserAddress)
			strictAuthRouter.DELETE("/user/address/:id", userAddressHandler.DeleteUserAddress)
		}
	}

	return s
}
