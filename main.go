package main

import (
	"net/http"
	"pet/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	
	
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
	r.Run()
}
