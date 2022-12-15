package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) UsersProxyIndex(c *gin.Context) {
	userProxy, err := h.services.UserProxy.FindById(4)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userProxy)

}

func (h *Handler) initUserProxyRoutes(api *gin.RouterGroup) {
	usersProxy := api.Group("/user-proxy")
	{
		usersProxy.GET("", h.UsersProxyIndex)
	}
}
