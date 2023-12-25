package middleware

import (
	"gin-bubble/common"
	"gin-bubble/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 终止当前请求,返回信息给客户端
			ctx.Abort()
			return
		}

		// 切割字符串=》Bearer 将token信息拿出来
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 终止当前请求,返回信息给客户端
			ctx.Abort()
			return
		}

		// 验证通过后获取claim 中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户
		if user.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 终止当前请求,返回信息给客户端
			ctx.Abort()
			return
		}

		// 用户存在 将user 的信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
