package router

import (
	"asocks-ws/internal/delivery/ws/v1/handlers"
	"asocks-ws/internal/service"
	"asocks-ws/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	services *service.Services
	apiToken string
}

func NewRouter(services *service.Services, apiToken string) *Router {
	return &Router{
		services: services,
		apiToken: apiToken,
	}
}

func (r *Router) Init() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.initAPI(router)
	return router
}

func (r *Router) initAPI(router *gin.Engine) {
	handlerV1 := handlers.NewHandler(r.services)
	api := router.Group("/api")
	api.Use(TokenAuthMiddleware("Bearer " + r.apiToken))
	{
		handlerV1.Init(api)
	}
}
func TokenAuthMiddleware(apiToken string) gin.HandlerFunc {
	if apiToken == "" {
		logger.Error("[Token not installed]")
	}
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")

		if token == "" {
			context.AbortWithStatusJSON(401, gin.H{"error": "API token required"})
			return
		}
		if token != apiToken {
			context.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
			return
		}
		context.Next()
	}
}
