package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DatabaseKey = "db"

func Database(c *gin.Context) {
	if username, usernameExists := os.LookupEnv("HARMONY_DB_USERNAME"); !usernameExists {
		// TODO: Log missing DB Username
	} else if password, passwordExists := os.LookupEnv("HARMONY_DB_PASSWORD"); !passwordExists {
		// TODO: Log missing DB Password
	} else if host, hostExists := os.LookupEnv("HARMONY_DB_HOST"); !hostExists {
		// TODO: Log missing DB Host
	} else if name, nameExists := os.LookupEnv("HARMONY_DB_NAME"); !nameExists {
		// TODO: Log missing DB Name
	} else if db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, name)), &gorm.Config{}); err != nil {
		// TODO: Log DB Connection Error
	} else {
		c.Set(DatabaseKey, db)
		c.Next()
		return
	}

	c.Status(http.StatusInternalServerError)
	c.Abort()
}
