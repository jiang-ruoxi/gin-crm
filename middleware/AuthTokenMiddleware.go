package middleware

import (
	"fmt"
	"gin-crm/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//AuthTokenMiddleware jwt中间件
func AuthTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//从header中获取Authorization的内容
		authorToken := ctx.GetHeader("Authorization")
		//判断authorToken是否存在，并且是以Bearer 开头的字符串
		if authorToken == "" || strings.HasSuffix(authorToken, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足，Authorization不存在或者格式有误"})
			ctx.Abort()
			return
		}
		//截取authorToken，从第8位到最后的内容
		authorToken = authorToken[7:]
		//校验token信息
		claims, err := lib.CheckTokenAuth(authorToken)
		fmt.Println(claims, err, claims.UserId)
		if err != nil || claims.UserId <= 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足，authorToken校验失败"})
			ctx.Abort()
			return
		}

		//验证通过后获取claim中的userId，并通过该用户id获取用户信息
		userId := claims.UserId
		//todo 这里是需要查询数据库新信息
		//
		//判断是否存在该用户的信息，如果不存在则进行拦截
		if userId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足，数据信息非法"})
			ctx.Abort()
			return
		}

		//将当前的用户信息写入上下文，需要将查询到的用户信息吧空字符串进行替换
		ctx.Set("login_user", "")

		ctx.Next()
	}
}
