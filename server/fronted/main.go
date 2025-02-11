package main

import (
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/haomiao000/Online-Shop/server/fronted/consts"
	"github.com/haomiao000/Online-Shop/server/fronted/router"
	"github.com/haomiao000/Online-Shop/server/fronted/bootstrap"
)

func main() {
	env := pflag.String("env", "dev", "指定运行环境: dev 或 online")
	pflag.Parse()
	if *env != "dev" && *env != "online" {
		log.Fatalf("启动参数错误: 仅支持 'dev' 或 'online'")
	}
	consts.GloablEnv = *env
	bootstrap.Init()
	address := viper.GetString("hertz.address")
	
	h := server.New(server.WithHostPorts(address))
	
	router.RegisterRoutes(h)
	h.Spin()
}