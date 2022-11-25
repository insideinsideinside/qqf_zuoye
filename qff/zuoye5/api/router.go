package api

import (
	"github.com/gin-gonic/gin"
	"zuoye3/api/middleware"
	"zuoye3/dao"
)

func Initrouter() {
	r := gin.Default()
	go dao.Readdatabase() //将文件读入假数据库
	r.Use(middleware.CORS())
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/changepassword", Changepassword) //实在不知道怎么本地文件操作完成删除数据，只能让最后一组user和password是真正的账号密码。。
	r.GET("/findpassword", FindPassword)
	r.POST("/writemessage", Message)
	r.GET("/checkmessage", Checkmessage)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}
	r.Run(":8088")
}
