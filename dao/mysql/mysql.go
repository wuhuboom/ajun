package mysql

import (
	"ajun/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() error {
	//連接數據庫
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	NameDatabase := viper.GetString("mysql.database")
	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/" + NameDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("连接成功")

	////////////////////////////////////////////////////////////////////
	//model的创建更新
	model.CheckIsExistModelUser(DB)
	////////////////////////////////////////////////////////////////////

	return nil
}
