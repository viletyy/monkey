/*
 * @Date: 2021-04-27 18:06:34
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 18:07:53
 * @FilePath: /monkey/routers/user.go
 */
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/controllers/user"
)

func init() {
	userNs := beego.NewNamespace("/user",
		beego.NSRouter("/", &user.User{}, "get:Index"),
		beego.NSRouter("/article", &user.Article{}, "get:Index;post:Create"),
		beego.NSRouter("/article/:id", &user.Article{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/article/new", &user.Article{}, "get:New"),
		beego.NSRouter("/article/:id/edit", &user.Article{}, "get:Edit"),
	)
	beego.AddNamespace(userNs)
}
