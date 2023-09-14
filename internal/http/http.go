package http

import (
	"net/http"
	"os"
	"strings"

	"github.com/NicholeMattera/Harmony/internal/app"
	"github.com/gin-gonic/gin"
)

type httpLayer struct {
    engine *gin.Engine
	app		app.AppLayer
}

type StandardObjectAuditing struct {
	CreatedBy     string `json:"createdBy"`
	CreatedOn     int64  `json:"createdOn"`
	DeletedBy     string `json:"deletedBy"`
	DeletedOn     int64  `json:"deletedOn"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	LastUpdatedOn int64  `json:"lastUpdatedOn"`
}

func New(appLayer app.AppLayer) *httpLayer {
    h := &httpLayer{
		engine: gin.New(),
        app: appLayer,
    }

	h.SetupFrontendRoutes()
	h.SetupApiRoutes()

	return h
}

func (h *httpLayer) SetupFrontendRoutes() {
	hostSite, hostSiteExists := os.LookupEnv("HARMONY_HOST_SITE")
	hostAdmin, hostAdminExists := os.LookupEnv("HARMONY_HOST_ADMIN")
	shouldHostSite := hostSiteExists && strings.ToLower(hostSite) == "true"
	shouldHostAdmin := hostAdminExists && strings.ToLower(hostAdmin) == "true"

	if shouldHostSite {
		h.engine.StaticFile("/index.html", "./assets/index.html")
		h.engine.StaticFile("/favicon.ico", "./assets/favicon.ico")
		h.engine.StaticFile("/robots.txt", "./assets/robots.txt")
	}

	if shouldHostAdmin {
		h.engine.StaticFile("/admin", "./assets/admin.html")
	}

	if shouldHostSite || shouldHostAdmin {
		h.engine.Static("/content", "./assets/content")
		h.engine.NoRoute(func(c *gin.Context) {
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
}

func (h *httpLayer) SetupApiRoutes() {
	api := h.engine.Group("/api/v1")
	{
		entity := api.Group("/entity")
		{
			entity.Use(Authorization)

			entity.GET("/", h.ListEntities)
			entity.GET("/slug/:slug", h.TakeEntity)
			entity.POST("/", h.CreateEntity)
			id := entity.Group("/:id")
			{
				id.GET("/", h.TakeEntity)
				id.PUT("/", h.UpdateEntity)
				id.DELETE("/", h.DeleteEntity)
			}
		}

		api.POST("/login", h.Login)

		role := api.Group("/role")
		{
			role.Use(Authorization)

			role.GET("/", h.ListRoles)
			role.POST("/", h.CreateRole)
			id := role.Group("/:id")
			{
				id.GET("/:id", h.TakeRole)
				id.PUT("/:id", h.UpdateRole)
				id.DELETE("/:id", h.DeleteRole)
			}
		}

		schema := api.Group("/schema")
		{
			schema.Use(Authorization)

			schema.GET("/", h.ListSchemas)
			schema.POST("/", h.CreateSchema)
			id := schema.Group("/:id")
			{
				id.GET("/", h.TakeSchema)
				id.PUT("/", h.UpdateSchema)
				id.DELETE("/", h.DeleteSchema)
			}
		}

		user := api.Group("/user")
		{
			user.Use(Authorization)

			user.GET("/", h.ListUsers)
			user.POST("/", h.CreateUser)
			id := user.Group("/:id")
			{
				id.GET("/", h.TakeUser)
				id.PUT("/", h.UpdateUser)
				id.DELETE("/", h.DeleteUser)
				id.DELETE("/sessions", h.InvalidateUserSessions)
			}
		}
	}
}

func (h *httpLayer) Run() {
	if addr, addrExists := os.LookupEnv("HARMONY_LISTEN_ADDR"); !addrExists {
		h.engine.Run("0.0.0.0:8080")
	} else {
		h.engine.Run(addr)
	}
}
