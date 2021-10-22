package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

func main() {
	debug := GetEnvInfo("MX_SHOP")
	var configName string
	if debug {
		configName = "config-dev.yaml"
	} else {
		configName = "config-prod.yaml"
	}
	v := viper.New()
	// 文件的路径，需要配置Working Directory之后运行才能拿到
	v.SetConfigFile(configName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	// 监听动态变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})
	time.Sleep(time.Second * 300)
}

// GetEnvInfo 根据环境变量读取不同的文件
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
