package middleware

import (
	"net/http"

	"github.com/NicholeMattera/Harmony/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var UserKey = "user"

func Authorization(c *gin.Context) {
	xsrf := c.GetHeader("X-Harmony-XSRF")
	if xsrf == "" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	db := c.MustGet(DatabaseKey).(*gorm.DB)

	user := model.User{Session: xsrf}
	result := db.First(&user)
	if result.Error != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	c.Set(UserKey, user)
	c.Next()
}
