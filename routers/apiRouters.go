package routers

import (
	"pet/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/allusers", api.Select)
		apiRouters.POST("/adduser", api.Add)
		apiRouters.PUT("/changeuser", api.Change)
		apiRouters.DELETE("/deluser", api.Del)
		apiRouters.POST("/register", api.Register)
	}
}
