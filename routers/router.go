/*
 * @Date: 2021-03-11 16:33:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:24:34
 * @FilePath: /monkey/routers/router.go
 */
package routers

import (
	"strings"
	"time"

	"github.com/viletyy/monkey/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, filterMethod)
	_ = beego.AddFuncMap("isActiveController", isActiveController)
	_ = beego.AddFuncMap("stringJoin", stringJoin)
	_ = beego.AddFuncMap("formatTime", formatTime)

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
	if strings.HasPrefix(parttern, controller) {
		return "is-active"
	}
	return ""
}

func stringJoin(s []string, sep string) string {
	return strings.Join(s, sep)
}

func filterMethod(ctx *context.Context) {
	if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
		ctx.Request.Method = strings.ToUpper(ctx.Input.Query("_method"))
	}
}

func formatTime(time time.Time) string {
	return time.Format("2006年1月2日 15:04:05")
}
