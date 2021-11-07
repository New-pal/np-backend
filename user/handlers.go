package user

import (
	"github.com/New-pal/np-backend/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userUpdate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    bool   `json:"gender"`
	Age       uint   `json:"age"`
}

type userSettingsUpdate struct {
	Language           string `json:"language"`
	LastNameVisibility bool   `json:"last_name_visibility"`
	AgeVisibility      bool   `json:"age_visibility"`
	TimeFormat         bool   `json:"time_format"`
	DistanceUnits      bool   `json:"distance_units"`
}

type Handler struct {
	ur  *UserRepository
	us  *UserService
	usr *UserSettingsRepository
	uss *UserSettingsService
}

// userGet - get current user
// @Summary get current user
// @Produce json
// @Success 200 {object} core.BaseApiResponse{results=[]user.User}
// @failure 401 {object} core.BaseApiErrorResponse
// @Router /api/user [get]
func (h *Handler) userGet(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		core.ApiErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	core.ApiResponse(c, http.StatusOK, u)
}

// userPatch - patch current user
// @Summary patch current user
// @Produce json
// @Param userUpdate body userUpdate true "userUpdate"
// @Success 200 {object} core.BaseApiResponse{results=[]user.User}
// @failure 401 {object} core.BaseApiErrorResponse
// @Router /api/user [patch]
func (h *Handler) userPatch(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		core.ApiErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err := c.ShouldBindJSON(u); err != nil {
		core.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if updateErr := h.us.UpdateUser(u); updateErr != nil {
		core.ApiErrorResponse(c, http.StatusInternalServerError, updateErr.Error())
		return
	}
	core.ApiResponse(c, http.StatusOK, u)
}

func NewUserHandler(ur *UserRepository, us *UserService, usr *UserSettingsRepository,
	uss *UserSettingsService) *Handler {
	return &Handler{ur: ur, us: us, usr: usr, uss: uss}
}
