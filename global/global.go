package global

import (
	"github.com/spf13/viper"
	"go-gin-chat/config"
	"gorm.io/gorm"
)

var (
	Viper  *viper.Viper
	Config *config.Config
	DB     *gorm.DB
)
