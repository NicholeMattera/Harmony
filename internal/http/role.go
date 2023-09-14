package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionBinding struct {
	CanRead  bool          `json:"canRead"`
	CanWrite bool          `json:"canWrite"`
	Schema   SchemaBinding `json:"schema"`
}

type RoleBinding struct {
	StandardObjectAuditing
	Id               int64               `json:"id"`
	LockToIp         bool                `json:"lockToIp"`
	MultipleSessions bool                `json:"multipleSessions"`
	Name             string              `json:"name"`
	SessionTimeout   int64               `json:"sessionTimeout"`
	Permissions      []PermissionBinding `json:"permissions"`
}

func (h *httpLayer) CreateRole(c *gin.Context) {
	var binding RoleBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}

func (h *httpLayer) DeleteRole(c *gin.Context) {

}

func (h *httpLayer) ListRoles(c *gin.Context) {

}

func (h *httpLayer) TakeRole(c *gin.Context) {

}

func (h *httpLayer) UpdateRole(c *gin.Context) {
	var binding RoleBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}
