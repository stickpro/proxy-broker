package router

import (
	"asocks-ws/internal/delivery/ws/v1/handlers"
	"asocks-ws/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	services *service.Services
}

func NewRouter(services *service.Services) *Router {
	return &Router{
		services: services,
	}
}

func (r *Router) Init() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.initAPI(router)
	return router
}

func (r *Router) initAPI(router *gin.Engine) {
	handlerV1 := handlers.NewHandler(r.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
