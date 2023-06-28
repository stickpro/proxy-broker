package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserUpdateForTS(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid or empty id query param")
		return
	}

	user, err := h.services.User.GetUserByID(ID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.User.SendUserToTSTopics(user); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
