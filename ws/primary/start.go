package primary

import (
	"github.com/gin-gonic/gin"
	"go-gin-chat/global"
	"go-gin-chat/ws"
	"go-gin-chat/ws/go_ws"
)

// 定义 serve 的映射关系
var serveMap = map[string]ws.ServeInterface{
	"Serve":   &ws.Serve{},
	"GoServe": &go_ws.GoServe{},
}

func Create() ws.ServeInterface {
	// GoServe or Serve
	_type := global.Config.App.ServeType
	return serveMap[_type]
}

func Start(gin *gin.Context) {
	Create().RunWs(gin)
}

func OnlineUserCount() int {
	return Create().GetOnlineUserCount()
}

func OnlineRoomUserCount(roomId int) int {
	return Create().GetOnlineRoomUserCount(roomId)
}
