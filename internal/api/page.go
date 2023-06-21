package api

import (
	"github.com/NicholeMattera/Harmony/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

type PageBinding struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PageHandler struct{}

func NewPageHandler(routerGroup *gin.RouterGroup) {
	pageHandler := &PageHandler{}

	pageRouterGroup := routerGroup.Group("/page")
	pageRouterGroup.GET("/", pageHandler.List)
	pageRouterGroup.GET("/:id", pageHandler.Take)
	pageRouterGroup.POST("/", pageHandler.Create).Use(middleware.Authorization)
	pageRouterGroup.PUT("/:id", pageHandler.Update).Use(middleware.Authorization)
	pageRouterGroup.DELETE("/:id", pageHandler.Delete).Use(middleware.Authorization)
}

func (*PageHandler) Create(c *gin.Context) {

}

func (*PageHandler) Delete(c *gin.Context) {

}

func (*PageHandler) List(c *gin.Context) {

}

func (*PageHandler) Take(c *gin.Context) {

}

func (*PageHandler) Update(c *gin.Context) {

}
