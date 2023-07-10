package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var UserKey = "user"

func Authorization(c *gin.Context) {
	xsrfHeaderName, xsrfHeaderNameExists := os.LookupEnv("HARMONY_XSRF_HEADER_NAME")
	if !xsrfHeaderNameExists {
		xsrfHeaderName = "X-Harmony-XSRF"
	}

	sessionCookieName, sessionCookieNameExists := os.LookupEnv("HARMONY_SESSION_COOKIE_NAME")
	if !sessionCookieNameExists {
		sessionCookieName = "harmony_session_token"
	}

	xsrfHeader := c.GetHeader(xsrfHeaderName)
	sessionCookie, err := c.Cookie(sessionCookieName)
	if err != nil || xsrfHeader == "" || sessionCookie == "" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	// TODO: Authorize with XSRF and Session Token

	c.Next()
}
