package http

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

func (h *httpLayer) CreateSchema(c *gin.Context) {
	var binding SchemaBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}

func (h *httpLayer) DeleteSchema(c *gin.Context) {

}

func (h *httpLayer) ListSchemas(c *gin.Context) {

}

func (h *httpLayer) TakeSchema(c *gin.Context) {

}

func (h *httpLayer) UpdateSchema(c *gin.Context) {
	var binding SchemaBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}
