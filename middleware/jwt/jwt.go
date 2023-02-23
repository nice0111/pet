package middleware

import (
	"fmt"
	"pet/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		// code = e.SUCCESS
		var token string
		if t := c.Request.Header.Get("Token"); t != "" {
			token = t
			// fmt.Println(t)
		} else {
			token = c.Query("Token")
		}

		// logging.Debug("Token:", token)
		if token == "" {
			// code = e.INVALID_PARAMS

			c.JSON(200, gin.H{
				"code": 409,
				"msg":  "409~~",
				"data": data,
			})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		fmt.Println("--------")
		// fmt.Printf("%v", err.Error())
		// fmt.Println(err)
		if err != nil {
			code = 400
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = 200
		}

		if code != 200 {
			c.JSON(200, gin.H{
				"code": 406,
				"msg":  "406",
				"data": data,
			})
			c.Abort()
			return
		}
		fmt.Printf("token-claims:%v\n", claims.Data.Id)
		c.Set("uid", claims.Data.Id)
		c.Next()
	}
}
