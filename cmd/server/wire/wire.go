//go:build wireinject
// +build wireinject

package wire

import (
	"aphrodite-go/internal/handler"
	"aphrodite-go/internal/repository"
	"aphrodite-go/internal/server"
	"aphrodite-go/internal/service"
	"aphrodite-go/pkg/app"
	"aphrodite-go/pkg/jwt"
	"aphrodite-go/pkg/log"
	"aphrodite-go/pkg/server/http"
	"aphrodite-go/pkg/sid"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewUserFeedbackRepository,
	repository.NewUserAddressRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewUserFeedbackService,
	service.NewUserAddressService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewAuthHandler,
	handler.NewUserHandler,
	handler.NewUserFeedbackHandler,
	handler.NewUserAddressHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("aphrodite-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
