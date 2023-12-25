package controller

import (
	"gin-bubble/model"
	"gin-bubble/response"
	"gin-bubble/service"
	"github.com/gin-gonic/gin"
	"time"
)

type BubbleController struct {
	BubbleService service.IBubbleService
}

// NewBubbleController 注册service层的数据内容,能够调用service中的接口方法
func NewBubbleController() BubbleController {
	bubbleController := BubbleController{BubbleService: service.NewBubbleService()}
	bubbleController.BubbleService.(service.BubbleService).DB.AutoMigrate(model.Todo{})
	return bubbleController
}

func (b BubbleController) AddBubble(ctx *gin.Context) {
	var todo model.Todo
	ctx.ShouldBind(&todo)
	// 设置新增时间--修改时间
	todo.Time = time.Now().Format("2006-01-02 15:04:05")
	// 调用服务层方法
	err := b.BubbleService.AddBubble(&todo)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil, "新增成功！")
}
func (b BubbleController) DelBubble(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")

	if !ok {
		response.Fail(ctx, "删除失败!", gin.H{"error": "无效的id"})
		return
	}
	err := b.BubbleService.DelBubble(id)
	if err != nil {
		response.Fail(ctx, err.Error(), gin.H{
			"error": "删除失败",
		})
		return
	}
	response.Success(ctx, nil, "删除成功!")
}
func (b BubbleController) UpdateBubble(ctx *gin.Context) {
	// 查询数据库内的记事本数据
	id, ok := ctx.Params.Get("id")
	if !ok {
		response.Fail(ctx, "更新失败!", gin.H{"error": "无效的id"})
		return
	}
	status, ok := ctx.Params.Get("status")
	if !ok {
		response.Fail(ctx, "更新失败!", gin.H{"error": "无效的id"})
		return
	}
	err := b.BubbleService.UpdateBubble(id, status)
	if err != nil {
		response.Fail(ctx, "更新失败!", gin.H{"error": err.Error()})
		return
	}
	response.Success(ctx, nil, "更新成功!")
}
func (b BubbleController) GetBubbleList(ctx *gin.Context) {
	// 查询数据库内的用户数据
	userId, ok := ctx.Params.Get("userId")
	if !ok {
		response.Fail(ctx, "查询失败!", gin.H{"error": "无效的id"})
		return
	}
	bubbleList, err := b.BubbleService.GetBubbleList(userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, gin.H{"bubbleList": bubbleList}, "查询成功")
}
