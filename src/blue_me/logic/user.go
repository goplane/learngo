package logic

import (
	"blue_me/dao/mysql"
	"blue_me/models"
	"blue_me/pkg/snowflake"
)

func SignUp(*models.ParamSignUp) {
	// 1.判断用户是否存在
	mysql.QueryUserByUsername()
	// 2.生成UID
	snowflake.GenID()
	// 3.密码加密
	// 4.保存进数据库
	mysql.InsertUser()
	// redis.XXX ...
}
