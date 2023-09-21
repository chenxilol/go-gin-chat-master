package session

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-gin-chat/global"
	"go-gin-chat/models"
	"net/http"
	"strconv"
)

func EnableCookieSession() gin.HandlerFunc {

	store := cookie.NewStore([]byte(global.Config.App.CookieKey))

	return sessions.Sessions("go-gin-chat", store)
}

// 注册和登陆时都需要保存seesion信息
func SaveAuthSession(c *gin.Context, info interface{}) {
	session := sessions.Default(c)
	session.Set("uid", info)
	// c.SetCookie("user_id",string(info.(map[string]interface{})["b"].(uint)), 1000, "/", "localhost", false, true)
	session.Save()
}

// GetSessionUserInfo 获取session中的用户信息
func GetSessionUserInfo(c *gin.Context) map[string]interface{} {
	session := sessions.Default(c)

	uid := session.Get("uid")

	data := make(map[string]interface{})
	if uid != nil {
		user := models.FindUserByField("id", uid.(string))
		data["uid"] = user.ID
		data["username"] = user.Username
		data["avatar_id"] = user.AvatarId
	}
	return data
}

// 退出时清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("uid"); sessionValue == nil {
		return false
	}
	return true
}

// AuthSessionMiddle 验证session中间件
func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("uid")
		fmt.Println(sessionValue)
		fmt.Println(c.Get("uid"))
		if sessionValue == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		uidInt, _ := strconv.Atoi(sessionValue.(string))

		if uidInt <= 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}

		// 设置简单的变量
		c.Set("uid", sessionValue)

		c.Next()
		return
	}
}
