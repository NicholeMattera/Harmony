package api

import (
	"github.com/NicholeMattera/Harmony/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

type UserBinding struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UserHandler struct{}

func NewUserHandler(routerGroup *gin.RouterGroup) {
	userHandler := &UserHandler{}

	userRouterGroup := routerGroup.Group("/user")
	userRouterGroup.GET("/", userHandler.List)
	userRouterGroup.GET("/:id", userHandler.Take)
	userRouterGroup.POST("/", userHandler.Create).Use(middleware.Authorization)
	userRouterGroup.PUT("/:id", userHandler.Update).Use(middleware.Authorization)
	userRouterGroup.DELETE("/:id", userHandler.Delete).Use(middleware.Authorization)
}

func (*UserHandler) Create(c *gin.Context) {

}

func (*UserHandler) Delete(c *gin.Context) {

}

func (*UserHandler) List(c *gin.Context) {

}

func (*UserHandler) Take(c *gin.Context) {

}

func (*UserHandler) Update(c *gin.Context) {

}
