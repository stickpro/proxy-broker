package service

import (
	"asocks-ws/internal/config"
	"asocks-ws/internal/repository"
)

type Services struct {
	UserProxy UserProxyServiceInterface
	Server    ServerServiceInterface
	User      UserServiceInterface
}
type Deps struct {
	Repository  *repository.Repositories
	KafkaConfig config.KafkaConfig
}

func NewServices(deps Deps) *Services {
	userProxyService := NewUserProxyService(deps.Repository.UserProxy, deps.KafkaConfig)
	serverService := NewServerService(deps.Repository.Server)
	userService := NewUserService(deps.Repository.User, deps.KafkaConfig)

	return &Services{
		UserProxy: userProxyService,
		Server:    serverService,
		User:      userService,
	}
}
