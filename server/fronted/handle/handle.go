package handle

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"context"
)

type AboutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}


type UIPresenter interface {
	Run() map[string]any
}


func(d AboutService) Run() map[string]any {
	return utils.H{
		"title": "About",
	}
}

//@router About
func About(ctx context.Context, c *app.RequestContext) {
	a := &AboutService{
		RequestContext: c,
		Context: ctx,
	}
}