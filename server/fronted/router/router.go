package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"context"
)

func RegisterRoutes(r *server.Hertz) {
	r.LoadHTMLGlob("template/*")
	r.Static("/static", "./")

	r.GET("/", func(c context.Context, ctx *app.RequestContext) {
		hlog.CtxInfof(c, "CloudeWeGo shop home page")
		ctx.HTML(200, "home", utils.H{"title": ""})
	})

	r.GET("/about", func(c context.Context, ctx *app.RequestContext) {
		hlog.CtxInfof(c, "CloudeWeGo shop about page")
		ctx.HTML(consts.StatusOK, "about", utils.H{"title": "About"})
	})


}