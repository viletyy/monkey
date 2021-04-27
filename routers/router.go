/*
 * @Date: 2021-03-11 16:33:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 22:45:59
 * @FilePath: /monkey/routers/router.go
 */
package routers

import (
	"strings"

	"github.com/viletyy/monkey/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, filterMethod)
	_ = beego.AddFuncMap("isActiveController", isActiveController)
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
}

func isActiveController(parttern, controller string) string {
	if parttern == controller {
		return "is-active"
	}
	return ""
}

func filterMethod(ctx *context.Context) {
	if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
		ctx.Request.Method = strings.ToUpper(ctx.Input.Query("_method"))
	}
}
