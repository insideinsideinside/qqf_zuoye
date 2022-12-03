package api

import (
	"github.com/gin-gonic/gin"
	"main/app/api/middleware"
	"main/app/internal/model"
)

func Initrouter() {
	r := gin.Default()
	model.InitDB()
	r.Use(middleware.CORS())
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/changepassword", Changepassword)
	r.GET("/findpassword", FindPassword)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}
	r.Run(":8088")
}
