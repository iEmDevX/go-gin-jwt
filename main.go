package main

import (
	"gin-jwt-em/auth"
	"gin-jwt-em/config"
	"gin-jwt-em/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	config.CommonConfig(r)

	// set route
	auth.SetRoute(r)
	user.SetRoute(r)

	if err := http.ListenAndServe(":"+config.Port, r); err != nil {
		log.Fatal(err)
	}
}
