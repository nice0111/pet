package middleware

import (
	"fmt"
	"pet/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var code int
		var data interface{}
		// code = e.SUCCESS
		var token string
		if t := c.Request.Header.Get("Token"); t != "" {
			token = t
			fmt.Println(t)
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

		// claims, err := utils.ParseToken(token)

		// fmt.Println(claims.StandardClaims.Issuer)
		// fmt.Println("--------")
		// fmt.Printf("%v", err.Error())
		// fmt.Println(err)
		// if err != nil {
		// 	code = 400
		// }
		// fmt.Println(time.Now().Unix())
		// fmt.Println(claims.StandardClaims.ExpiresAt)
		// code = 200
		// if time.Now().Unix() > claims.ExpiresAt {
		// 	fmt.Println("token已过期")
		// 	code = 200
		// }
		// 	code = 200
		// }

		// if code != 200 {
		// 	c.JSON(200, gin.H{
		// 		"code": 406,
		// 		"msg":  "406",
		// 		"data": data,
		// 	})
		// 	c.Abort()
		// 	return
		// }
		// // fmt.Printf("token-claims:%v\n", claims.Data.Id)
		// c.Set("uid", claims.Data.Id)
		// // c.Set("uid", claims.(*utils.Claims).Data.ExpiresAt)
		// c.Next()

		// claims := new(jwt.Claims)
		tokens, _ := jwt.ParseWithClaims(token, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("aaaaaaaaaaaaaaaa"), nil
		})
		claims, ok := tokens.Claims.(*utils.Claims)
		if !ok {
			c.JSON(401, gin.H{
				"msg": "失败",
			})
			c.Abort()
			return
		}
		fmt.Printf("token-claims:%v\n", claims.Data.Id)
		c.Next()
	}
}
