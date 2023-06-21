package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginBinding struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginHandler struct{}

func NewLoginHandler(router *gin.Engine) {
	loginHandler := &LoginHandler{}

	router.POST("/login", loginHandler.Login)
}

func (*LoginHandler) Login(c *gin.Context) {
	var binding LoginBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}
