/*
 * @Date: 2021-04-23 10:06:42
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 21:25:17
 * @FilePath: /monkey/controllers/admin/base.go
 */
package admin

import (
	"html/template"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/utils"
	"github.com/viletyy/yolk"
)

type Base struct {
	beego.Controller
	FlashBag         *beego.FlashData
	RedirectUrl      string
	CurrentLoginUser *model.User
	StartTime        int64
	HandlerSeconds   float64
}

func (c *Base) Finish() {
	handlerSecond := float64(time.Now().UnixNano()-c.StartTime) / float64(1e9)
	c.HandlerSeconds = handlerSecond
}

func (c *Base) Prepare() {
	c.TplPrefix = "admin/"
	c.Layout = "layout/admin.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navbar"] = "shared/navbar.tpl"
	c.LayoutSections["Notification"] = "shared/notification.tpl"
	c.LayoutSections["Menu"] = "shared/admin_menu.tpl"
	c.LayoutSections["Footer"] = "shared/footer.tpl"

	// 启动时间
	c.StartTime = time.Now().UnixNano()

	// 初始化读取Flash
	beego.ReadFromRequest(&c.Controller)

	// 初始化Flash
	c.FlashBag = beego.NewFlash()

	// 初始化XSRF
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	// 自动读取当前登陆用户
	isLogin := false
	isAdmin := false
	loginUserId, _ := strconv.Atoi(c.Ctx.GetCookie("login_user_id"))
	loginUserSign := c.Ctx.GetCookie("login_user_sign")
	if loginUserId > 0 {
		user, err := model.GetUserById(loginUserId)
		if err == nil && utils.AuthSignCheck(int(user.ID), user.Username, user.Password, loginUserSign) {
			c.CurrentLoginUser = &user
			isLogin = true
			isAdmin = user.IsAdmin
		}
	}
	c.Data["User"] = c.CurrentLoginUser
	c.Data["IsLogin"] = isLogin
	c.Data["IsAdmin"] = isAdmin

	if !isAdmin {
		c.RedirectTo("/user")
	}
}

// 保存成功的Flash信息
func (c *Base) FlashSuccess(message string) {
	c.FlashBag.Notice(message)
	c.FlashBag.Store(&c.Controller)
}

// 保存失败的Flash信息
func (c *Base) FlashError(message string) {
	c.FlashBag.Error(message)
	c.FlashBag.Store(&c.Controller)
}

// 自动化的表单验证器
func (c *Base) ValidatorAuto(data interface{}) {
	if err := c.ParseForm(data); err != nil {
		c.FlashBag.Error("参数解析错误")
		c.RedirectTo(c.RedirectUrl)
	}

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
	validate := validation.Validation{}

	isValid, err := validate.Valid(data)
	if err != nil {
		c.ErrorHandler(err)
	}

	if !isValid {
		c.FlashError(validate.Errors[0].Message)
		c.RedirectTo(c.RedirectUrl)
	}
}

// 重定向
func (c *Base) RedirectTo(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

func (c *Base) ErrorHandler(err error) {
	logs.Info(err)
	c.Abort("500")
	c.StopRun()
}

// 跳转到前一页
func (c *Base) Back() {
	c.RedirectTo(c.Ctx.Request.Referer())
	c.StopRun()
}

// 默认分页参数
func (c *Base) PageInfo() (pageInfo global.PageInfo) {
	var err error
	pageInfo.Page, err = c.GetInt("page")
	if pageInfo.Page <= 0 || err != nil {
		pageInfo.Page = 1
	}
	pageInfo.PageSize, err = c.GetInt("per")
	if pageInfo.PageSize <= 0 || pageInfo.PageSize > 15 || err != nil {
		pageInfo.PageSize = 15
	}
	return
}

func (c *Base) ResponseWithResult(searchResult global.SearchResult) {
	p := yolk.NewPaginator(c.Ctx.Request, searchResult.PageSize, searchResult.Total)
	c.Data["paginator"] = p
	c.Data["List"] = searchResult.List
}
