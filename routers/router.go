/*
 * @Date: 2021-03-11 16:33:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-25 22:55:14
 * @FilePath: /monkey/routers/router.go
 */
package routers

import (
	"github.com/viletyy/monkey/controllers"
	"github.com/viletyy/monkey/controllers/admin"
	"github.com/viletyy/monkey/controllers/user"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.Index{}, "get:Index")
	beego.Router("/search", &controllers.Index{}, "get:Search")
	beego.Router("/article", &controllers.Article{}, "get:Index")
	beego.Router("/article/:id", &controllers.Article{}, "get:Show")
	beego.Router("/news", &controllers.News{}, "get:Index")
	beego.Router("/news/:id", &controllers.News{}, "get:Show")
	beego.Router("/plate", &controllers.Plate{}, "get:Index")
	beego.Router("/plate/:id", &controllers.Plate{}, "get:Show")
	beego.Router("/login", &controllers.User{}, "get:Login;post:LoginHandle")
	beego.Router("/register", &controllers.User{}, "get:Register;post:RegisterHandle")
	beego.Router("/loginout", &controllers.User{}, "get:Loginout")

	userNs := beego.NewNamespace("/user",
		beego.NSRouter("/", &user.User{}, "get:Index"),
		beego.NSRouter("/article", &user.Article{}, "get:Index;post:Create"),
		beego.NSRouter("/article/:id", &user.Article{}, "put:Update;delete:Destroy"),
		beego.NSRouter("/article/new", &user.Article{}, "get:New"),
		beego.NSRouter("/article/:id/edit", &user.Article{}, "get:Edit"),
	)
	beego.AddNamespace(userNs)

	adminNs := beego.NewNamespace("/admin",
		beego.NSRouter("/", &admin.Dashboard{}, "get:Index"),
		beego.NSRouter("/plate", &admin.Plate{}, "get:Index"),
		beego.NSRouter("/article", &admin.Article{}, "get:Index"),
		beego.NSRouter("/news", &admin.News{}, "get:Index"),
		beego.NSRouter("/tag", &admin.Tag{}, "get:Index"),
		beego.NSRouter("/user", &admin.User{}, "get:Index"),
		beego.NSRouter("/setting", &admin.Setting{}, "get:Index"),
		beego.NSRouter("/banner", &admin.Banner{}, "get:Index"),
		beego.NSRouter("/recommend", &admin.Recommend{}, "get:Index"),
	)

	beego.AddNamespace(adminNs)
	_ = beego.AddFuncMap("isActiveController", isActiveController)
}

func isActiveController(parttern, controller string) string {
	if parttern == controller {
		return "is-active"
	}
	return ""
}
