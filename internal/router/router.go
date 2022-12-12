package router

import (
	"asocks-ws/internal/delivery/ws/v1/handlers"
	"asocks-ws/internal/service"
)

type Router struct {
	services *service.Services
}

func NewRouter(services *service.Services) *Router {
	return &Router{
		services: services,
	}
}

func (r *Router) Init() {
	handlerV1 := handlers.NewHandler(r.services)
	handlerV1.Init()
}
