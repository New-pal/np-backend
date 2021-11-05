package auth

import "github.com/gin-gonic/gin"

func BindRouter(group *gin.RouterGroup, handler *Handler) {
	group.POST("/login", handler.login)
	group.POST("/register", handler.register)
	group.POST("/refresh", handler.refresh)
	group.POST("/logout", handler.logout)
}
