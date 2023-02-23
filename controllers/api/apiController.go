package api

import (
	"fmt"
	"net/http"
	"pet/models"
	"pet/utils"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	// "gorm.io/gorm/utils"
)

// 查询所有用户
func Select(ctx *gin.Context) {
	usersList := []models.User{}
	models.DB.Find(&usersList)
	ctx.JSON(http.StatusOK, gin.H{
		"result": usersList,
	})
	// ctx.String(200, "ok")
}

// 新增用户
func Add(ctx *gin.Context) {
	phone := ctx.Query("phone")
	fmt.Println(phone)
	user := models.User{Phone: phone}
	res := models.DB.First(&user, "phone = ?", phone)
	// result.RowsAffected  返回找到的记录数
	// result.Error         returns error or nil
	fmt.Println(res.RowsAffected)
	if res.RowsAffected == 0 {
		models.DB.Create(&user)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "新建成功",
			"data": &user,
		})
	} else {
		ctx.JSON(403, gin.H{
			"msg":    "创建失败",
			"resson": phone + "用户已存在",
		})
	}
}

// 删除用户
func Del(ctx *gin.Context) {
	phone := ctx.Query("phone")
	user := models.User{Phone: phone}
	res := models.DB.First(&user, "phone = ?", phone)
	if res.RowsAffected != 0 {
		models.DB.Where("phone = ?", phone).Delete(&user)
		ctx.JSON(200, gin.H{
			"Code": 200,
			"msg":  "删除成功",
		})
	} else {
		ctx.JSON(200, gin.H{
			"Code": 300,
			"msg":  "数据不存在",
		})
	}

	// ctx.String(200, "Del")
}

// 更新用户
func Change(ctx *gin.Context) {
	// phone := ctx.Query("phone")
	// user := models.User{Phone: phone}
	// res := models.DB.First(&user, "phone = ?", phone)
	// if res.RowsAffected != 0 {
	// 	models.DB.Where("phone = ?", phone).Update(&user)
	// 	ctx.JSON(200, gin.H{
	// 		"Code": 200,
	// 		"msg":  "更新成功",
	// 	})
	// } else {
	// 	ctx.JSON(200, gin.H{
	// 		"Code": 300,
	// 		"msg":  "数据不存在",
	// 	})
	// }
	ctx.String(200, "Change")
}

// 用户注册
func Register(ctx *gin.Context) {
	phone := ctx.DefaultPostForm("phone", "")
	username := ctx.DefaultPostForm("username", "")
	// passwd := ctx.DefaultPostForm("123", "")

	vaild := validation.Validation{}
	vaild.Phone(phone, "phone").Message("请输入手机号")
	// vaild.Phone(passwd, "passwd").Message("请输入密码")
	if vaild.HasErrors() {
		ctx.JSON(500, gin.H{
			"msg": vaild.Errors,
		})
	} else {
		var jwttoken utils.JwtData
		jwttoken.Id = 1
		// jwttoken.Username = username
		// jwttoken.Mobile = phone

		token, _ := utils.GenerateToken(jwttoken)
		user := models.User{
			Phone:     phone,
			Username:  username,
			Password:  ctx.DefaultPostForm("password", ""),
			Email:     ctx.DefaultPostForm("email", ""),
			Gender:    ctx.DefaultPostForm("gender", ""),
			Logintime: time.Now().Unix(),
			Loginip:   ctx.ClientIP(),
			City:      ctx.DefaultPostForm("city", ""),
			Token:     token,
		}
		res := models.DB.First(&user, "phone = ?", phone)
		if res.RowsAffected == 0 {
			models.DB.Create(&user)
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "注册成功",
				"data": &user,
			})
		} else {
			ctx.JSON(403, gin.H{
				"msg":    "注册失败",
				"resson": phone + "用户已存在",
			})
		}
	}

}
