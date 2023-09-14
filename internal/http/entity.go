package http

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

func (h *httpLayer) CreateEntity(c *gin.Context) {
	var binding EntityBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}

func (h *httpLayer) DeleteEntity(c *gin.Context) {

}

func (h *httpLayer) ListEntities(c *gin.Context) {

}

func (h *httpLayer) TakeEntity(c *gin.Context) {

}

func (h *httpLayer) UpdateEntity(c *gin.Context) {
	var binding EntityBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}
