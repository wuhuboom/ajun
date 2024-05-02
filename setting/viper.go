package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

// Init 初始化
func Init() error {
	viper.SetConfigName("config") // 配置文件名 (无扩展名)
	viper.AddConfigPath(".")      // 可选的配置文件路径
	err := viper.ReadInConfig()   // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return err
	}
	return nil
}
