package user

import (
	"ajun/dao/mysql"
	"ajun/model"
	"ajun/tools"
	"github.com/gin-gonic/gin"
	"time"
)

//注册

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//1.判斷用戶名和密碼是否為空?
	if password == "" || username == "" {
		tools.ReturnResponse(c, tools.ErrorBackCode, nil, "賬號或密碼不能為空!")
		return
	}
	//2.判斷數據庫裡面是否存在這個數據
	user := model.User{}
	result := mysql.DB.Where("username = ?", username).First(&user)
	if result.RowsAffected > 0 {
		tools.ReturnResponse(c, tools.ErrorBackCode, nil, "用戶已存在")
		return
	}
	//对密码进行加密处理
	password = tools.MD5hash(password)
	//3.生成唯一憑證  58為隨機字符串  0-9 a-z A-Z  還需要一個判定是否重讀
	user.Token = string(tools.GenerateRandomString(58)) // 生成新的 token
	//4.判斷沒有問題入庫
	user.Username = username
	user.Created = time.Now().Unix()
	err := mysql.DB.Save(&user).Error
	if err != nil {
		tools.ReturnResponse(c, tools.ErrorBackCode, nil, err.Error())
		return
	}
	tools.ReturnResponse(c, tools.SuccessBackCode, nil, "註冊成功")
	return
}

//登录
