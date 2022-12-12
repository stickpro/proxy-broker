package handlers

import "asocks-ws/internal/service"

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() {
	h.initUserProxyRoutes()
}
