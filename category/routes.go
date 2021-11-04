package category

import "github.com/gin-gonic/gin"

func BindRouter(router *gin.Engine, handler *Handler) {
	group := router.Group("/api")
	{
		group.GET("/categories", handler.categoryList)
		group.GET("/categories/:id", handler.categoryGet)
		group.GET("/subcategories", handler.subcategoryList)
		group.GET("/subcategories/:id", handler.subcategoryGet)
	}
}
