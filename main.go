package main

import (
	"ajun/dao/mysql"
	"ajun/router"
	"ajun/setting"
	"fmt"
)

func main() {
	//1.读取配置文件
	err := setting.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 2.连接数据库
	err = mysql.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//开启gin框架
	router.Init()
}
