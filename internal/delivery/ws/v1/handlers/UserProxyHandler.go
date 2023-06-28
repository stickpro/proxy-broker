package handlers

import (
	"asocks-ws/pkg/logger"
	"net/http"

	"asocks-ws/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UsersProxyIndex(c *gin.Context) {
	server, err := h.services.Server.FindByIP("45.82.65.183")
	if err != nil {
		return
	}
	userProxies, err := h.services.UserProxy.FindByServerIP(server.Ip)
	if err != nil {
		return
	}
	go func() {
		h.services.UserProxy.SendKafkaInitMessage(userProxies, "test")
	}()

	//todo есть ли смысл этого действия, если ошибка сюда не доходит?
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userProxies)
}

func (h *Handler) UserProxyUpdateForLpm(c *gin.Context) {
	var request domain.UserProxy

	if err := c.BindJSON(&request); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("[UserProxyHandler request]", request)
	userProxy, err := h.services.UserProxy.FindById(request.Id)
	logger.Info("[UserProxyHandler UserProxy]", userProxy)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// todo remove this spike
	userProxy.Status = request.Status

	_, err = h.services.UserProxy.SendKafkaMessage(userProxy, request.IP)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Error send broker")
		return
	}

	c.JSON(http.StatusOK, userProxy)
}

func (h *Handler) initUserProxyRoutes(api *gin.RouterGroup) {
	usersProxy := api.Group("/user-proxy")
	{
		usersProxy.GET("", h.UsersProxyIndex)
		usersProxy.POST("/lpm-update", h.UserProxyUpdateForLpm)
	}
}
