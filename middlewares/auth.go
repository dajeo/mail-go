package middlewares

import (
	"github.com/gin-gonic/gin"
	"mail/config"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatus(401)
		}

		cfg := config.GetConfig()

		req, errReq := http.NewRequest(
			http.MethodPost,
			cfg.GetString("wp.url")+"jwt-auth/v1/token/validate",
			nil,
		)
		req.Header.Add("Authorization", token)
		if errReq != nil {
			c.AbortWithStatus(500)
		}

		res, errRes := http.DefaultClient.Do(req)
		if errRes != nil {
			c.AbortWithStatus(500)
		}

		if res.StatusCode != 200 {
			c.AbortWithStatus(401)
		}

		c.Next()
	}
}
