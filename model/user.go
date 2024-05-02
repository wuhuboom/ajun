package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID           int     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	Token        string  `json:"token"`                                       //token 唯一憑證
	FundPassword string  `json:"fund_password"`                               // 支付密碼
	Balance      float64 `gorm:"type:decimal(10,2);default:0" json:"balance"` //餘額
	Created      int64   `json:"created"`                                     //註冊時間
	Updated      int64   `json:"updated"`
}

func CheckIsExistModelUser(db *gorm.DB) {
	if db.Migrator().HasTable(&User{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&User{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.Migrator().CreateTable(&User{})
	}
}
