package controller

import (
	"gin-bubble/common"
	"gin-bubble/dto"
	"gin-bubble/model"
	"gin-bubble/response"
	"gin-bubble/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// UserController 才能调用controller下的方法 且注册service参数
type UserController struct {
	UserService service.IUserService
}

// NewUserController 注册service层的数据内容,能够调用service中的接口方法
func NewUserController() UserController {
	userController := UserController{UserService: service.NewUserService()}
	userController.UserService.(service.UserService).DB.AutoMigrate(model.User{})
	return userController
}

// Login 登录
func (u UserController) Login(ctx *gin.Context) {
	var loginUser dto.UserDto
	ctx.Bind(&loginUser)

	// 获取参数
	name := loginUser.Name
	number := loginUser.Number
	password := loginUser.Password

	// 数据验证
	if len(number) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 调用方法
	user, err := u.UserService.Login(name, number)

	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	// 校验密码
	if user.Password != password {
		response.Fail(ctx, "密码不正确,请重新输入", nil)
		return
	}

	//  校验通过,发放token
	token, err := common.ReleaseToken(*user)
	if err != nil {
		response.Response(ctx, 500, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

// UserInfo 获取用户个人信息
func (u UserController) UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	if user == nil {
		response.Fail(ctx, "用户不存在", nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
