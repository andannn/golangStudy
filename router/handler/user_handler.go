package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Hello(context *gin.Context) {
	user, _ := h.service.GetUserById(1)

	context.JSON(200, user)
}
