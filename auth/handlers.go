package auth

import (
	"github.com/New-pal/np-backend/core"
	"github.com/New-pal/np-backend/user"
	gwt "github.com/ennaque/go-gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	ur           userRepositoryInterface
	us           userServiceInterface
	authHandlers *gwt.Handler
}

// register - registration endpoint
// @Summary Registers user
// @Produce json
// @Param RegisterCredentials body RegisterCredentials true "RegisterCredentials"
// @Success 200 {object} core.BaseApiResponse{results=[]user.User}
// @failure 400 {object} core.BaseApiErrorResponse
// @failure 500 {object} core.BaseApiErrorResponse
// @Router /auth/register [post]
func (h *Handler) register(c *gin.Context) {
	var credentials RegisterCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		core.ApiErrorResponse(c, http.StatusBadRequest, ErrInvalidCredentials.Error())
		return
	}
	usr, err := h.us.CreateUser(credentials.Email, credentials.Name, credentials.Password)
	if err != nil {
		core.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	core.ApiResponse(c, http.StatusOK, usr.(*user.User))
}

// login - login endpoint
// @Summary Login user
// @Produce json
// @Param LoginCredentials body LoginCredentials true "LoginCredentials"
// @Success 200 {object} gwt.DefaultLoginResponse
// @failure 400 {object} gwt.DefaultErrResponse
// @failure 401 {object} gwt.DefaultErrResponse
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	h.authHandlers.GetLoginHandler()(c)
}

// refresh - refresh endpoint
// @Summary Refresh tokens
// @Produce json
// @Param gwt.RefreshRequestData body gwt.RefreshRequestData true "gwt.RefreshRequestData"
// @Success 200 {object} gwt.DefaultLoginResponse
// @failure 400 {object} gwt.DefaultErrResponse
// @failure 401 {object} gwt.DefaultErrResponse
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	h.authHandlers.GetRefreshHandler()(c)
}

// logout - logout endpoint
// @Summary Logout user
// @Produce json
// @Success 200 {object} gwt.DefaultLogoutResponse
// @failure 401 {object} gwt.DefaultErrResponse
// @Router /auth/logout [post]
func (h *Handler) logout(c *gin.Context) {
	h.authHandlers.GetLogoutHandler()(c)
}

func NewAuthHandler(ur interface{ userRepositoryInterface }, us interface{ userServiceInterface }, auth *gwt.Handler) *Handler {
	return &Handler{ur: ur, us: us, authHandlers: auth}
}
