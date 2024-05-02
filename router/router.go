package router

import (
	ControllerUser "ajun/controller/user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	r := gin.Default()
	user := r.Group("user/api", ControllerUser.Register)
	{
		//1.注册
		user.POST("register")
	}

	r.Run(":" + viper.GetString("project.port"))
}
