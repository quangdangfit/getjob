package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/jwt"
	"github.com/quangdangfit/gocommon/logger"
	"go.uber.org/dig"

	"github.com/quangdangfit/getjob/pkg/errors"
	"github.com/quangdangfit/getjob/pkg/http/wrapper"
)

// JWT middleware constants
const (
	UserKey = "user_data"
)

// GetToken get token from Header of gin.Context
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// JWT middleware handle authenticate by jwt
func JWT(container *dig.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := container.Invoke(func(
			auth jwt.IJWTAuth,
		) {
			token := GetToken(c)
			if token == "" {
				err := errors.New(errors.ECInvalidCredentials, "Unauthorized")
				c.Header("WWW-Authenticate", "jwt realm=\"gin\"")
				c.JSON(http.StatusUnauthorized, wrapper.Response{Error: err})
				logger.Errorf("Unauthorized, token: %s : %s", token, err)

				c.Abort()
				return
			}

			payload, err := auth.ValidateToken(token)
			if err != nil {
				c.Header("WWW-Authenticate", "jwt realm=\"gin\"")
				c.JSON(http.StatusUnauthorized, wrapper.Response{Error: err})
				logger.Errorf("Unauthorized, token: %s : %s", token, err)

				c.Abort()
				return
			}

			c.Set("id", payload[UserKey])
			c.Next()
		})

		if err != nil {
			logger.Errorf("Failed to invoke objects: %s", err)
		}
	}
}
