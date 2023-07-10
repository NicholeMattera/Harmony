package api

import (
	"github.com/NicholeMattera/Harmony/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

type StandardObjectAuditing struct {
	CreatedBy     string `json:"createdBy"`
	CreatedOn     int64  `json:"createdOn"`
	DeletedBy     string `json:"deletedBy"`
	DeletedOn     int64  `json:"deletedOn"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	LastUpdatedOn int64  `json:"lastUpdatedOn"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Database)
	router.Use(middleware.Authorization)

	router.StaticFile("/", "./assets/site.html")
	router.StaticFile("/admin", "./assets/admin.html")
	router.Static("/content", "./assets/content/")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.StaticFile("/robots.txt", "./assets/robots.txt")

	NewLoginHandler(router)

	apiV1RouterGroup := router.Group("/api/v1")
	NewEntityHandler(apiV1RouterGroup)
	NewRoleHandler(apiV1RouterGroup)
	NewSchemaHandler(apiV1RouterGroup)
	NewUserHandler(apiV1RouterGroup)

	return router
}
