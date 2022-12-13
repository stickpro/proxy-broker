package service

import (
	"asocks-ws/internal/repository"
)

type Services struct {
	UserProxy UserProxyServiceInterface
}
type Deps struct {
	Repository *repository.Repositories
}

func NewServices(deps Deps) *Services {
	userProxyService := NewUserProxyService(deps.Repository.UserProxy)
	return &Services{
		UserProxy: userProxyService,
	}
}
