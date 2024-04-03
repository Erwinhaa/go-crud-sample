package middlewares

import (
	"context"
	"log"
	"myapp/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		bearer := "Bearer "

		if auth == "" || len(auth) <= len(bearer) {
			c.Next()
			return
		}

		auth = auth[len(bearer):]

		validate, err := tools.JwtValidate(auth)
		if err != nil || !validate.Valid {
			log.Println(err)
			c.Next()
			return
		}

		claim, _ := validate.Claims.(*tools.JwtCustomClaim)

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), tools.JwtCtxKey, claim))

		c.Next()
	}
}

func IsUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if tools.AuthCtx(c.Request.Context()) == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
