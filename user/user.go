package user

import (
	"gin-jwt-em/auth"
	"gin-jwt-em/constants"
	"gin-jwt-em/model"

	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
)

func SetRoute(r *gin.Engine) {
	user := r.Group("/user")
	user.Use(auth.GetMiddleware().MiddlewareFunc())
	{
		user.POST("/refresh_token", auth.GetMiddleware().RefreshHandler)
		user.GET("/hello", helloHandler)
	}
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(constants.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[constants.IdentityKey],
		"userName": user.(*model.User).UserName,
		"text":     "Hello World.",
	})
}
