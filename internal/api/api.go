package api

import (
	"net/http"
	"os"
	"strings"

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

	// Frontend Routes
	hostSite, hostSiteExists := os.LookupEnv("HARMONY_HOST_SITE")
	hostAdmin, hostAdminExists := os.LookupEnv("HARMONY_HOST_ADMIN")
	shouldHostSite := hostSiteExists && strings.ToLower(hostSite) == "true"
	shouldHostAdmin := hostAdminExists && strings.ToLower(hostAdmin) == "true"

	if shouldHostSite {
		router.StaticFile("/index.html", "./assets/index.html")
		router.StaticFile("/favicon.ico", "./assets/favicon.ico")
		router.StaticFile("/robots.txt", "./assets/robots.txt")
	}

	if shouldHostAdmin {
		router.StaticFile("/admin", "./assets/admin.html")
	}

	if shouldHostSite || shouldHostAdmin {
		router.Static("/content", "./assets/content")
		router.NoRoute(func(c *gin.Context) {
			if shouldHostAdmin && len(c.Request.URL.Path) >= 5 && c.Request.URL.Path[:5] == "admin" {
				c.Header("Content-Type", "text/html; charset=utf-8")
				c.File("./dist/admin.html")
			} else if shouldHostSite {
				c.Header("Content-Type", "text/html; charset=utf-8")
				c.File("./dist/index.html")
			}

			c.Status(http.StatusNotFound)
		})
	}

	// Login Routes
	NewLoginHandler(router)

	// API Routes
	apiV1RouterGroup := router.Group("/api/v1")
	NewEntityHandler(apiV1RouterGroup)
	NewRoleHandler(apiV1RouterGroup)
	NewSchemaHandler(apiV1RouterGroup)
	NewUserHandler(apiV1RouterGroup)

	return router
}
