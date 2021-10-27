package auth

import (
	"github.com/New-pal/np-backend/core"
	"github.com/gin-gonic/gin"
)

func Authenticator(ur interface{ userRepositoryInterface }) func(c *gin.Context) (string, error) {
	return func(c *gin.Context) (string, error) {
		var credentials LoginCredentials
		if err := c.ShouldBindJSON(&credentials); err != nil {
			return "", ErrInvalidCredentials
		}
		usr, usrErr := ur.GetUserByEmail(credentials.Email)
		if usrErr != nil {
			return "", ErrInvalidCredentials
		}
		if passErr := core.CheckPassword(usr.GetPassword(), credentials.Password); passErr != nil {
			return "", ErrInvalidCredentials
		}
		return usr.GetId(), nil
	}
}
