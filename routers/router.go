/*
 * @Date: 2021-03-11 16:33:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-22 18:36:26
 * @FilePath: /monkey/routers/router.go
 */
package routers

import (
	"github.com/viletyy/monkey/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/search", &controllers.IndexController{}, "get:Search")
	beego.Router("/article", &controllers.ArticleController{}, "get:Index")
	beego.Router("/news", &controllers.NewsController{}, "get:Index")
	beego.Router("/user", &controllers.UserController{}, "get:Index")
	beego.Router("/login", &controllers.UserController{}, "get:Login;post:LoginHandle")
	beego.Router("/register", &controllers.UserController{}, "get:Register;post:RegisterHandle")
	beego.Router("/loginout", &controllers.UserController{}, "get:Loginout")
}
