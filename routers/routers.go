package routers

import (
	"gin-bubble/controller"
	"gin-bubble/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	// 配置跨越中间件以及监测程序是否异常的中间件
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	userController := controller.NewUserController()
	// 登录接口
	r.POST("/login", userController.Login)
	// 获取用户信息
	r.Use(middleware.AuthMiddleware()).GET("/info", userController.UserInfo)

	bubbleController := controller.NewBubbleController()
	todoRouters := r.Group("/todo")
	// 配置token验证
	todoRouters.Use(middleware.AuthMiddleware())
	{
		// 获取用户待办事项列表
		todoRouters.GET("/list/:userId", bubbleController.GetBubbleList)
		// 增加待办事项
		todoRouters.POST("/add", bubbleController.AddBubble)
		// 删除用户
		todoRouters.DELETE("/:id", bubbleController.DelBubble)
		// 修改
		todoRouters.PUT("/:id/:status", bubbleController.UpdateBubble)
	}
	return r
}
