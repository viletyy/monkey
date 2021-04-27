/*
 * @Date: 2021-04-27 18:06:28
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 18:12:20
 * @FilePath: /monkey/routers/admin.go
 */
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/controllers/admin"
)

func init() {
	adminNs := beego.NewNamespace("/admin",
		beego.NSRouter("/", &admin.Dashboard{}, "get:Index"),
		beego.NSRouter("/plate", &admin.Plate{}, "get:Index"),
		beego.NSRouter("/article", &admin.Article{}, "get:Index"),
		beego.NSRouter("/news", &admin.News{}, "get:Index"),
		beego.NSRouter("/tag", &admin.Tag{}, "get:Index"),
		// 用户管理
		beego.NSRouter("/user", &admin.User{}, "get:Index;post:Create"),
		beego.NSRouter("/user/new", &admin.User{}, "get:New"),
		beego.NSRouter("/user/:id", &admin.User{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/user/:id/edit", &admin.User{}, "get:Edit"),
		beego.NSRouter("/user/:id/reset", &admin.User{}, "put:Reset"),

		beego.NSRouter("/setting", &admin.Setting{}, "get:Index"),
		beego.NSRouter("/banner", &admin.Banner{}, "get:Index"),
		beego.NSRouter("/recommend", &admin.Recommend{}, "get:Index"),
	)

	beego.AddNamespace(adminNs)
}
