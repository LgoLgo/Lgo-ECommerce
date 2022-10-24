package initialize

import (
	middlewares "E-commerce-system/apis/user-web/middleware"
	userRouter "E-commerce-system/apis/user-web/router"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Routers(port int) *server.Hertz {
	Router := server.Default(
		server.WithHostPorts(fmt.Sprintf(":%d", port)),
	)

	//配置跨域
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/v1")
	userRouter.InitUserRouter(ApiGroup)

	return Router
}
