package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FieldValueBinding struct {
	Field FieldBinding `json:"field"`
	Value string       `json:"value"`
}

type EntityBinding struct {
	StandardObjectAuditing
	FieldValues    []FieldValueBinding `json:"fieldValues"`
	Id             int64               `json:"id"`
	Name           string              `json:"name"`
	ParentEntityId int64               `json:"parentEntityId"`
	Schema         SchemaBinding       `json:"schema"`
	Slug           string              `json:"slug"`
}

type EntityHandler struct{}

func NewEntityHandler(routerGroup *gin.RouterGroup) {
	entityHandler := &EntityHandler{}

	entityRouterGroup := routerGroup.Group("/entity")
	entityRouterGroup.GET("/", entityHandler.List)
	entityRouterGroup.GET("/:id", entityHandler.Take)
	entityRouterGroup.GET("/slug/:slug", entityHandler.Take)
	entityRouterGroup.POST("/", entityHandler.Create)
	entityRouterGroup.PUT("/:id", entityHandler.Update)
	entityRouterGroup.DELETE("/:id", entityHandler.Delete)
}

func (*EntityHandler) Create(c *gin.Context) {
	var binding EntityBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}

func (*EntityHandler) Delete(c *gin.Context) {

}

func (*EntityHandler) List(c *gin.Context) {

}

func (*EntityHandler) Take(c *gin.Context) {

}

func (*EntityHandler) Update(c *gin.Context) {
	var binding EntityBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}
