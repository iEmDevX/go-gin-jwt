package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Port : port server
var Port string

// CommonConfig : Common Config app
func CommonConfig(r *gin.Engine) {
	Port = os.Getenv("PORT")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if Port == "" {
		Port = "8000"
	}
}
