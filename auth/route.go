package auth

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) {
	r.POST("/login",GetMiddleware().LoginHandler)

	r.NoRoute(GetMiddleware().MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code":  "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}

