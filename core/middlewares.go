package core

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func CorsMiddleware() gin.HandlerFunc {
	//corsConfig := cors.DefaultConfig()
	////corsConfig.AllowAllOrigins = true
	//corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	//corsMiddleware := cors.New(corsConfig)
	return cors.Default()
}
