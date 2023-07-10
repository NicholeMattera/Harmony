package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserBinding struct {
	StandardObjectAuditing
	Id        int64  `json:"id"`
	Email	  string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	Username  string `json:"username"`
}

type UserHandler struct{}

func NewUserHandler(routerGroup *gin.RouterGroup) {
	userHandler := &UserHandler{}

	userRouterGroup := routerGroup.Group("/user")
	userRouterGroup.GET("/", userHandler.List)
	userRouterGroup.GET("/:id", userHandler.Take)
	userRouterGroup.POST("/", userHandler.Create)
	userRouterGroup.PUT("/:id", userHandler.Update)
	userRouterGroup.DELETE("/:id", userHandler.Delete)
	userRouterGroup.DELETE("/:id/sessions", userHandler.Invalidate)
}

func (*UserHandler) Create(c *gin.Context) {
	var binding UserBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}

func (*UserHandler) Delete(c *gin.Context) {

}

func (*UserHandler) List(c *gin.Context) {

}

func (*UserHandler) Take(c *gin.Context) {

}

func (*UserHandler) Update(c *gin.Context) {
	var binding UserBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}

func (*UserHandler) Invalidate(c *gin.Context) {

}
