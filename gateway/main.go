package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	// "github.com/haomiao000/Online-Shop/gateway/router"
)

func main() {
	h := server.Default()
	
	// router.RegisterRoutes(h)

	h.Spin()
}