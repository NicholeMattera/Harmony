package api

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

type RoleHandler struct{}

func NewRoleHandler(routerGroup *gin.RouterGroup) {
	roleHandler := &RoleHandler{}

	roleRouterGroup := routerGroup.Group("/role")
	roleRouterGroup.GET("/", roleHandler.List)
	roleRouterGroup.GET("/:id", roleHandler.Take)
	roleRouterGroup.POST("/", roleHandler.Create)
	roleRouterGroup.PUT("/:id", roleHandler.Update)
	roleRouterGroup.DELETE("/:id", roleHandler.Delete)
}

func (*RoleHandler) Create(c *gin.Context) {
	var binding RoleBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}

func (*RoleHandler) Delete(c *gin.Context) {

}

func (*RoleHandler) List(c *gin.Context) {

}

func (*RoleHandler) Take(c *gin.Context) {

}

func (*RoleHandler) Update(c *gin.Context) {
	var binding RoleBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}
