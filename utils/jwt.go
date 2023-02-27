package utils

import (

	// "heshang_go/pkg/setting"

	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("aaaaaaaaaaaaaaaa")

type Claims struct {
	// Username string `json:"username"`
	// Password string `json:"password"`
	Data JwtData `json:"data"`
	jwt.StandardClaims
}

type JwtData struct {
	Id int `json:"id"`
	// Username string `json:"username"`
	// Nickname string `json:"nickname"`
	// Mobile string `json:"mobile"`
}

func GenerateToken(data JwtData) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(72 * time.Hour)
	claims := Claims{
		data,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "heshang_go",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// fmt.Printf("jwt-tokenClaims:%v\n", tokenClaims)
	if err != nil {
		fmt.Println("有错误")
	}
	if tokenClaims != nil {
		fmt.Println("正在执行")
		fmt.Printf("jwt-tokenClaims:%v\n", tokenClaims.Claims.(*Claims))
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			fmt.Println("-------~~~~")
			fmt.Printf("jwt-claims:%v\n", claims)
			fmt.Printf("%v", &claims)
			return claims, nil

		}

	}
	return nil, err
}
