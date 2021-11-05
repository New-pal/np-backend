package user

import "github.com/gin-gonic/gin"

func BindRouter(group *gin.RouterGroup, handler *Handler) {
	group.GET("/user", handler.userGet)
	group.PATCH("/user", handler.userPatch)
}
