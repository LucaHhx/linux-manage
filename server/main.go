package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/websocket"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"net/http"
	"strings"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.6.2
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	core.Init()
	if global.GVA_DB != nil && global.GVA_CONFIG.Mysql.Migrate {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	go RunWebsocket()
	core.RunWindowsServer()
}

func RunWebsocket() {
	components := &component.Components{}
	components.Register(
		global.GVA_SOCKET,
		component.WithName("socket"), // rewrite component and handler name
		component.WithNameFunc(strings.ToLower),
	)

	nano.Listen(":3250",
		nano.WithIsWebsocket(true),
		nano.WithPipeline(global.GVA_SOCKET.Pip),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
		nano.WithWSPath("/socket"),
		nano.WithDebugMode(),
		nano.WithSerializer(websocket.NewSerializer()), // override default serializer
		nano.WithComponents(components),
	)
}
