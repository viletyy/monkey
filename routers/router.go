/*
 * @Date: 2021-03-11 16:33:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-23 13:59:33
 * @FilePath: /monkey/routers/router.go
 */
package routers

import (
	"github.com/viletyy/monkey/controllers"
	"github.com/viletyy/monkey/controllers/admin"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.Index{}, "get:Index")
	beego.Router("/search", &controllers.Index{}, "get:Search")
	beego.Router("/article", &controllers.Article{}, "get:Index")
	beego.Router("/news", &controllers.News{}, "get:Index")
	beego.Router("/user", &controllers.User{}, "get:Index")
	beego.Router("/login", &controllers.User{}, "get:Login;post:LoginHandle")
	beego.Router("/register", &controllers.User{}, "get:Register;post:RegisterHandle")
	beego.Router("/loginout", &controllers.User{}, "get:Loginout")

	adminNs := beego.NewNamespace("/admin",
		beego.NSRouter("/", &admin.Dashboard{}, "get:Index"),
		beego.NSRouter("/article", &admin.Article{}, "get:Index"),
	)

	beego.AddNamespace(adminNs)
}
