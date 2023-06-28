package handlers

import (
	"asocks-ws/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

/* http://arlimus.github.io/articles/gin.and.gorilla/  just try*/
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUserProxyRoutes(v1)
	}

	// POST /api/v1/users?id={int}
	v1.POST("/users", h.UserUpdateForTS)
}
