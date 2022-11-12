package initialize

import (
	"apis/oss-web/middlewares"
	"apis/oss-web/router"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func Routers(port int) *server.Hertz {
	Router := server.Default(server.WithHostPorts(fmt.Sprintf(":%d", port)))
	Router.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, utils.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	Router.LoadHTMLFiles(fmt.Sprintf("oss-web/templates/index.html"))
	Router.StaticFS("/static", &app.FS{Root: "oss-web/static"})
	Router.GET("", func(ctx context.Context, c *app.RequestContext) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "index.html", utils.H{
			"title": "posts/index",
		})
	})

	//配置跨域
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/oss/v1")
	router.InitOssRouter(ApiGroup)

	return Router
}
