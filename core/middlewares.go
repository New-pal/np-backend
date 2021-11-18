package core

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func CorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = viper.GetStringSlice("cors.white_list")
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE", "PUT"}
	corsMiddleware := cors.New(corsConfig)
	return corsMiddleware
}
