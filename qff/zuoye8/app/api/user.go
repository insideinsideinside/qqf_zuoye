package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/app/api/middleware"
	"main/app/internal/model"
	"main/utils"
	"time"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.ResponseFail(c, "verification failed")
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	flag := model.SelectUser(username)
	if flag {
		utils.ResponseFail(c, "user already exists")
		return
	}
	model.AddUser(username, password)
	utils.ResponseSuccess(c, "register success")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.ResponseFail(c, "verification failed")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := model.SelectUser(username)
	if !flag {
		utils.ResponseFail(c, "user doesn't exists")
		return
	}
	selectpassword := model.SelectPassword(username)
	if selectpassword != password {
		utils.ResponseFail(c, "wrong password")
		return
	}
	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Wzy",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.ResponseSuccess(c, tokenString)
}
func Changepassword(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	newpassword := c.PostForm("newpassword")
	flag := model.SelectUser(username)
	if !flag {
		utils.ResponseFail(c, "user doesn't exists")
		return
	}
	if password != model.SelectPassword(username) {
		utils.ResponseFail(c, "wrong password")
		return
	}
	u := model.User{
		Username: username,
		Password: newpassword,
		ID:       model.SelectID(username),
	}
	model.ChangePassword(u)
	flag1 := model.SelectUser(username)
	if flag1 {
		utils.ResponseSuccess(c, "password change success")
		return
	}
}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.ResponseSuccess(c, username.(string))
}

func FindPassword(c *gin.Context) {
	username := c.PostForm("username")
	flag := model.SelectUser(username)
	if !flag {
		utils.ResponseFail(c, "user doesn't exists")
		return
	}
	utils.ResponseSuccess(c, model.SelectPassword(username))
}
