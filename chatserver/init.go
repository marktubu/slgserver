package chatserver

import (
	"slgserver/chatserver/controller"
	"slgserver/net"
)

var MyRouter = &net.Router{}

func Init() {
	//全部初始化完才注册路由，防止服务器还没启动就绪收到请求
	initRouter()
}

func initRouter() {
	controller.DefaultChat.InitRouter(MyRouter)
}
