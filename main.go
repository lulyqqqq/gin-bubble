package main

import (
	"fmt"
	"gin-bubble/common"
	"gin-bubble/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// 程序入口文件
func main() {
	// 初始化连接信息
	InitConfig()
	// 连接mysql
	db := common.InitMySQL()
	fmt.Println(db)
	r := gin.Default()

	// 注册路由
	r = routers.SetupRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

// InitConfig 获取应用文件信息
func InitConfig() {
	// 获取工作目录
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
