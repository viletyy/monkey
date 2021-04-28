/*
 * @Date: 2021-04-27 18:06:28
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 23:39:33
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
		// 文章管理
		beego.NSRouter("/article", &admin.Article{}, "get:Index;post:Create"),
		beego.NSRouter("/article/new", &admin.Article{}, "get:New"),
		beego.NSRouter("/article/:id", &admin.Article{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/article/:id/edit", &admin.Article{}, "get:Edit"),
		// 资讯管理
		beego.NSRouter("/news", &admin.News{}, "get:Index;post:Create"),
		beego.NSRouter("/news/new", &admin.News{}, "get:New"),
		beego.NSRouter("/news/:id", &admin.News{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/news/:id/edit", &admin.News{}, "get:Edit"),
		// 板块管理
		beego.NSRouter("/plate", &admin.Plate{}, "get:Index;post:Create"),
		beego.NSRouter("/plate/new", &admin.Plate{}, "get:New"),
		beego.NSRouter("/plate/:id", &admin.Plate{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/plate/:id/edit", &admin.Plate{}, "get:Edit"),
		// 标签管理
		beego.NSRouter("/tag", &admin.Tag{}, "get:Index;post:Create"),
		beego.NSRouter("/tag/new", &admin.Tag{}, "get:New"),
		beego.NSRouter("/tag/:id", &admin.Tag{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/tag/:id/edit", &admin.Tag{}, "get:Edit"),
		// 用户管理
		beego.NSRouter("/user", &admin.User{}, "get:Index;post:Create"),
		beego.NSRouter("/user/new", &admin.User{}, "get:New"),
		beego.NSRouter("/user/:id", &admin.User{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/user/:id/edit", &admin.User{}, "get:Edit"),
		beego.NSRouter("/user/:id/reset", &admin.User{}, "put:Reset"),
		// 设置
		beego.NSRouter("/setting", &admin.Setting{}, "get:Index;post:Create"),
		beego.NSRouter("/setting/new", &admin.Setting{}, "get:New"),
		beego.NSRouter("/setting/:id", &admin.Setting{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/setting/:id/edit", &admin.Setting{}, "get:Edit"),
		// banner
		beego.NSRouter("/banner", &admin.Banner{}, "get:Index;post:Create"),
		beego.NSRouter("/banner/new", &admin.Banner{}, "get:New"),
		beego.NSRouter("/banner/:id", &admin.Banner{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/banner/:id/edit", &admin.Banner{}, "get:Edit"),
		// 推荐
		beego.NSRouter("/recommend", &admin.Recommend{}, "get:Index;post:Create"),
		beego.NSRouter("/recommend/new", &admin.Recommend{}, "get:New"),
		beego.NSRouter("/recommend/:id", &admin.Recommend{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/recommend/:id/edit", &admin.Recommend{}, "get:Edit"),
	)

	beego.AddNamespace(adminNs)
}
