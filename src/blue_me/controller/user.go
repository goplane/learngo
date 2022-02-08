package controller

import (
	"blue_me/logic"
	"blue_me/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	// 内置的new函数创建一个值类型
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应,可以记录个日志zap
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求的参数有误",
		})
		return
	}
	//上面的shouldbindjson只能简单的记录输入信息，并不能做限制，所以需要手动对请求参数进行详细的业务规则校验
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
		zap.L().Error("SignUp with invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求的参数有误",
		})
		return
	}

	fmt.Println(p)
	// 2.业务处理

	logic.SignUp(p)
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
