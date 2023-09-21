package models

import (
	"go-gin-chat/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username" json:"username" binding:"required,max=16,min=2"`
	Password string `form:"password" json:"password" binding:"required,max=32,min=6"`
	AvatarId string `form:"avatar_id" json:"avatar_id" binding:"required,numeric"`
}

// AddUser 添加用户
func AddUser(value interface{}) User {
	var u User
	u.Username = value.(map[string]interface{})["username"].(string)
	u.Password = value.(map[string]interface{})["password"].(string)
	u.AvatarId = value.(map[string]interface{})["avatar_id"].(string)
	global.DB.Create(&u)
	return u
}

// SaveAvatarId  保存头像
func SaveAvatarId(AvatarId string, u User) User {
	u.AvatarId = AvatarId
	global.DB.Save(&u)
	return u
}

// FindUserByField 根据字段查找用户
func FindUserByField(field, value string) User {
	var u User

	if field == "id" || field == "username" {
		global.DB.Where(field+" = ?", value).First(&u)
	}

	return u
}

func GetOnlineUserList(uids []float64) []map[string]interface{} {
	var results []map[string]interface{}
	global.DB.Where("id IN ?", uids).Find(&results)

	return results
}
