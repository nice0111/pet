package routers

import (
	"pet/controllers/admin"

	// "github.com/dgrijalva/jwt-go"
	middleware "pet/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/admin", middleware.JWT())
	{
		apiRouters.GET("/index", admin.Index)
	}
}
