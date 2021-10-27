package core

import (
	"github.com/gin-gonic/gin"
)

type BaseApiResponse struct {
	Count   int           `json:"count"`
	Results []interface{} `json:"results"`
}

type BaseApiErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func ApiResponse(c *gin.Context, code int, entities ...interface{}) {
	c.JSON(code, BaseApiResponse{Count: len(entities), Results: entities})
}

func ApiErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, BaseApiErrorResponse{ErrorCode: code, ErrorMessage: message})
	c.Abort()
}
