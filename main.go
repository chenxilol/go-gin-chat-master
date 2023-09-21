package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-chat/core"
	"go-gin-chat/global"
	"go-gin-chat/routes"
	"go-gin-chat/views"
	"go-gin-chat/ws/go_ws"
	"log"
)

func main() {
	// 关闭debug模式
	global.Viper = core.Viper()
	global.DB = core.InitGorm()
	gin.SetMode(gin.ReleaseMode)
	port := global.Config.App.Port
	router := routes.InitRoute()
	//加载模板文件
	router.SetHTMLTemplate(views.GoTpl)
	go_ws.CleanOfflineConn()

	log.Println("监听端口", "http://127.0.0.1:"+port)
	router.Run(":" + port)
	//http.ListenAndServe(":"+port, router)
}
