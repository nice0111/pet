package main

import (
	"net/http"
	"pet/routers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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

// func main() {
// 	// 加密串
// 	mySigningKey := []byte("flyfly123")

// 	c := &MyClaims{
// 		Username: "feifei",
// 		StandardClaims: jwt.StandardClaims{
// 			NotBefore: time.Now().Unix() - 60,
// 			ExpiresAt: time.Now().Unix() + 60*60*2,
// 			Issuer:    "feifei",
// 		},
// 	}
// 	// 创建token
// 	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
// 	fmt.Println(t)
// 	s, e := t.SignedString(mySigningKey)
// 	if e != nil {
// 		fmt.Printf("%s", e)
// 	}
// 	println(s)

// 	// 解析token
	// token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
	// 	return mySigningKey, nil
	// })
// 	fmt.Println(err)
// 	fmt.Println(token.Claims.(*MyClaims))
// 	fmt.Println(token.Claims.(*MyClaims).Username)

// }
