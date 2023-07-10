package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FieldBinding struct {
	StandardObjectAuditing
	DefaultValue string `json:"defaultValue"`
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Order        uint8  `json:"order"`
	Required     bool   `json:"required"`
	Type         string `json:"type"`

}

type SchemaBinding struct {
	StandardObjectAuditing
	CanReadByDefault  bool           `json:"canReadByDefault"`
	CanWriteByDefault bool           `json:"canWriteByDefault"`
	Id                int64          `json:"id"`
	Name              string         `json:"name"`
	Fields            []FieldBinding `json:"fields"`
}

type SchemaHandler struct{}

func NewSchemaHandler(routerGroup *gin.RouterGroup) {
	schemaHandler := &SchemaHandler{}

	schemaRouterGroup := routerGroup.Group("/schema")
	schemaRouterGroup.GET("/", schemaHandler.List)
	schemaRouterGroup.GET("/:id", schemaHandler.Take)
	schemaRouterGroup.POST("/", schemaHandler.Create)
	schemaRouterGroup.PUT("/:id", schemaHandler.Update)
	schemaRouterGroup.DELETE("/:id", schemaHandler.Delete)
}

func (*SchemaHandler) Create(c *gin.Context) {
	var binding SchemaBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}

func (*SchemaHandler) Delete(c *gin.Context) {

}

func (*SchemaHandler) List(c *gin.Context) {

}

func (*SchemaHandler) Take(c *gin.Context) {

}

func (*SchemaHandler) Update(c *gin.Context) {
	var binding SchemaBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}
