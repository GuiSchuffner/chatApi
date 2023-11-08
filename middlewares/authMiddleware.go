package middlewares

import (
	"net/http"

	"github.com/GuiSchuffner/chatApi/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.VerifyToken(c.GetHeader("Authorization"))
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
