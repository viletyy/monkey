/*
 * @Date: 2021-03-09 09:57:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 14:29:26
 * @FilePath: /monkey/main.go
 */
package main

import (
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/initialize"
	_ "github.com/viletyy/monkey/routers"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	global.LOG = initialize.Zap()

	// 初始化数据库
	global.DB = initialize.Gorm()
	defer global.DB.Close()

	// 初始化redis
	global.REDIS = initialize.Redis()
	defer global.REDIS.Close()

	// 设置validation信息

	defaultMessage := map[string]string{
		"Required":     "不能为空",
		"Min":          "不能小于%d",
		"Max":          "不能大于%d",
		"Range":        "取值必须在%d到%d之间",
		"MinSize":      "长度不能小于%d",
		"MaxSize":      "长度不能大于%d",
		"Length":       "长度必须等于%d",
		"Alpha":        "必须是字母",
		"Numeric":      "必须是数字",
		"AlphaNumeric": "必须是字母或者数字",
		"Match":        "必须出现 %s 关键字",
		"NoMatch":      "不能出现 %s 关键字",
		"AlphaDash":    "必须是字母，数组或者横线(-)",
		"Email":        "不合法的邮箱地址",
		"IP":           "不合法的IP",
		"Base64":       "不合法的Base64编码格式",
		"Mobile":       "不合法的手机号",
		"Tel":          "不合法的电话号码",
		"Phone":        "不合法的手机号",
		"ZipCode":      "不合法的邮编",
	}
	validation.SetDefaultMessage(defaultMessage)

	// 启动应用
	beego.Run()
}
