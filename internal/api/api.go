package api

import (
	"github.com/NicholeMattera/Harmony/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Database)

	router.StaticFile("/", "./assets/site.html")
	router.StaticFile("/admin", "./assets/admin.html")
	router.Static("/content", "./assets/content/")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.StaticFile("/robots.txt", "./assets/robots.txt")

	NewLoginHandler(router)

	apiV1RouterGroup := router.Group("/api/v1")
	NewPageHandler(apiV1RouterGroup)
	NewUserHandler(apiV1RouterGroup)

	return router
}
