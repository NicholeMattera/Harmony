package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginBinding struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *httpLayer) Login(c *gin.Context) {
	var binding LoginBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}
