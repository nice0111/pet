package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
	// ctx.String(200, "ok")
}
