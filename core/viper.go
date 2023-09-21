package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-gin-chat/core/internal"
	"go-gin-chat/global"
	"os"
)

func Viper(path ...string) *viper.Viper {
	// 优先级: 命令行 > 环境变量 > 默认值
	// 判断命令行参数是否为空
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您在使用的是%s环境变量,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您在使用的是%s环境变量,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您在使用的是%s环境变量,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				default:
					config = internal.ConfigDefaultFile
					fmt.Printf("您在使用的是%s环境变量,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				}
			} else {
				config = configEnv
				fmt.Printf("您在使用的是%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("您在使用的是-c参数传递的值,config的路径为%s\n", config)
		}
	} else {
		config = path[0]
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("被修改的配置文件名为%s", in.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println("序列化错误")
		}
	})
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println("序列化错误")
	}
	return v
}
