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
	//conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	//
	//for {
	//	msgType, msg, err := conn.ReadMessage()
	//	if err != nil {
	//		return
	//	}
	//	fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
	//	if err = conn.WriteMessage(msgType, msg); err != nil {
	//		return
	//	}
	//}
	data := struct {
		Name string
		Age  int
	}{"Joh Doe", 30}
	c.JSON(http.StatusOK, data)

}

func (h *Handler) initUserProxyRoutes(api *gin.RouterGroup) {
	usersProxy := api.Group("/users")
	{
		usersProxy.GET("", h.UsersProxyIndex)
	}
}
