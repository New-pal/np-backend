package category

import (
	"github.com/New-pal/np-backend/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	cr *CategoryRepository
	cs *CategoryService
	sr *SubcategoryRepository
	ss *SubCategoryService
}

// categoryGetList - get categories list
// @Summary get categories list
// @Produce json
// @Success 200 {object} core.BaseApiResponse{results=[]category.Category}
// @failure 500 {object} core.BaseApiErrorResponse
// @Router /api/categories [get]
func (h *Handler) categoryList(c *gin.Context) {
	var data []*Category
	var err error
	data, err = h.cr.getAll()
	if err != nil {
		core.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	core.ApiResponse(c, http.StatusOK, data)
}

// categoryGetList - get category by id
// @Summary get category by id
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} core.BaseApiResponse{results=[]category.Category}
// @failure 500 {object} core.BaseApiErrorResponse
// @Router /api/categories/{id} [get]
func (h *Handler) categoryGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		core.ApiErrorResponse(c, http.StatusBadRequest, ErrIdIsNotProvided.Error())
		return
	}
	category, e := h.cr.getById(id)
	if e != nil {
		core.ApiErrorResponse(c, http.StatusNotFound, e.Error())
		return
	}
	core.ApiResponse(c, http.StatusOK, category)
}

// subcategoryGetList - get subcategories list
// @Summary get subcategories list
// @Produce json
// @Success 200 {object} core.BaseApiResponse{results=[]category.Subcategory}
// @failure 500 {object} core.BaseApiErrorResponse
// @Router /api/subcategories [get]
func (h *Handler) subcategoryList(c *gin.Context) {
	var data []*Subcategory
	var err error
	data, err = h.sr.getAll()
	if err != nil {
		core.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	core.ApiResponse(c, http.StatusOK, data)
}

// categoryGetList - get subcategory by id
// @Summary get subcategory by id
// @Produce json
// @Param id path int true "Subcategory ID"
// @Success 200 {object} core.BaseApiResponse{results=[]category.Subcategory}
// @failure 500 {object} core.BaseApiErrorResponse
// @Router /api/subcategories/{id} [get]
func (h *Handler) subcategoryGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		core.ApiErrorResponse(c, http.StatusBadRequest, ErrIdIsNotProvided.Error())
		return
	}
	subcategory, e := h.sr.getById(id)
	if e != nil {
		core.ApiErrorResponse(c, http.StatusNotFound, e.Error())
		return
	}
	core.ApiResponse(c, http.StatusOK, subcategory)
}

func NewCategoryHandler(cr *CategoryRepository, cs *CategoryService,
	sr *SubcategoryRepository, ss *SubCategoryService) *Handler {
	return &Handler{cr: cr, cs: cs, sr: sr, ss: ss}
}
