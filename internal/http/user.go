package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserBinding struct {
	StandardObjectAuditing
	Id        int64  `json:"id"`
	Email	  string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	Username  string `json:"username"`
}

func (*httpLayer) CreateUser(c *gin.Context) {
	var binding UserBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}

func (*httpLayer) DeleteUser(c *gin.Context) {

}

func (*httpLayer) ListUsers(c *gin.Context) {

}

func (*httpLayer) TakeUser(c *gin.Context) {

}

func (*httpLayer) UpdateUser(c *gin.Context) {
	var binding UserBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

}

func (*httpLayer) InvalidateUserSessions(c *gin.Context) {

}
