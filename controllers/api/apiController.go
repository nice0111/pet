package api

import (
	"fmt"
	"net/http"
	"pet/models"
	"pet/utils"
	"strconv"
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

	vaild := validation.Validation{}
	vaild.Phone(phone, "phone").Message("请输入手机号")
	if vaild.HasErrors() {
		ctx.JSON(500, gin.H{
			"msg": vaild.Errors,
		})
	} else {
		var jwttoken utils.JwtData
		jwttoken.Id = 1

		user := models.User{
			Phone:     phone,
			Username:  username,
			Password:  ctx.DefaultPostForm("password", ""),
			Email:     ctx.DefaultPostForm("email", ""),
			Gender:    ctx.DefaultPostForm("gender", ""),
			Logintime: time.Now().Unix(),
			Loginip:   ctx.ClientIP(),
			City:      ctx.DefaultPostForm("city", ""),
			// Token:     token,
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

// 用户登录
func Login(ctx *gin.Context) {
	phone := ctx.DefaultPostForm("phone", "")
	password := ctx.DefaultPostForm("password", "")

	user := new(models.User)
	user.Phone = phone
	user.Password = password

	res := models.DB.First(&user, "phone = ?", phone)
	if res.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"code": 404,
			"msg":  "用户尚未注册，请先注册",
		})
	} else if password == (*user).Password {
		var jwttoken utils.JwtData
		jwttoken.Id = 1
		token, _ := utils.GenerateToken(jwttoken)
		// 用户登录,返回token
		ctx.JSON(200, gin.H{
			"code":  200,
			"msg":   "登录成功",
			"token": token,
		})
		// 用户每次登录,都更新user表的token值
		models.DB.Model(&user).Where("phone = ?", (*user).Phone).Update("token", token)
		fmt.Print("--------")
	} else {
		// 用户名密码不对的情况
		ctx.JSON(400, gin.H{
			"code": 401,
			"msg":  "登录失败",
			// "token": token,
		})
	}

}

// 传入宠物标识，返回宠物列表
func PetId(ctx *gin.Context) {
	sid := ctx.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(sid)
	petsList := []models.PetsName{}
	res := models.DB.Where("petid = ?", id).Find(&petsList)
	names := make(map[int]string)
	for i, v := range petsList {
		fmt.Println(i+1, v.Name)
		names[i+1] = v.Name
	}
	if res.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"code": 404,
			"msg":  "您查找的内容不存在",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  names,

			//  petsList[i].Name,
		})
	}
}

// 查询热门宠物的id
func Ishot(ctx *gin.Context) {
	sid := ctx.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(sid)
	petsList := []models.PetsName{}
	res := models.DB.Where("petid = ? AND ishot = ?", id, 1).Find(&petsList)
	if res.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"code": 404,
			"msg":  "信息未找到",
		})
	} else {
		names := make(map[int]string)
		for i, v := range petsList {
			fmt.Println(i+1, v.Name)
			names[i+1] = v.Name
		}
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  names,
		})
	}

}
