package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NicholeMattera/Harmony/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DatabaseKey = "db"

func initAndMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Session{})
	db.AutoMigrate(&model.Composition{})
	db.AutoMigrate(&model.Permission{})
	db.AutoMigrate(&model.Field{})
	db.AutoMigrate(&model.Entity{})
	db.AutoMigrate(&model.FieldValue{})
}

func Database() gin.HandlerFunc {
	username, usernameExists := os.LookupEnv("HARMONY_DB_USERNAME")
	if !usernameExists {
		log.Fatal("Error: No database username provided")
		return func(c *gin.Context) {
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}
	}

	password, passwordExists := os.LookupEnv("HARMONY_DB_PASSWORD")
	if !passwordExists {
		log.Fatal("Error: No database password provided")
		return func(c *gin.Context) {
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}
	}

	host, hostExists := os.LookupEnv("HARMONY_DB_HOST")
	if !hostExists {
		log.Fatal("Error: No database hostname provided")
		return func(c *gin.Context) {
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}
	}

	name, nameExists := os.LookupEnv("HARMONY_DB_NAME")
	if !nameExists {
		log.Fatal("Error: No database name provided")
		return func(c *gin.Context) {
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, name)), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: Unable to connect to MySQL server", err)
	}

	initAndMigrate(db)

	return func(c *gin.Context) {
		c.Set(DatabaseKey, db)
		c.Next()
	}
}
