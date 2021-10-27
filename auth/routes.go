package auth

import "github.com/gin-gonic/gin"

func BindRouter(router *gin.Engine, handler *Handler) {
	group := router.Group("/auth")
	{
		group.POST("/login", handler.login)
		group.POST("/register", handler.register)
		group.POST("/refresh", handler.refresh)
		group.POST("/logout", handler.logout)
	}
}
