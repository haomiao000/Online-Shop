package initializer

import (
	"github.com/spf13/viper"
	"github.com/haomiao000/Online-Shop/server/fronted/consts"
	"log"
)

func InitConf() {
	viper.SetConfigName(consts.GloablEnv)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("配置文件读取失败, %s", err)
	}
}