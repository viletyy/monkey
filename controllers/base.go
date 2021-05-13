/*
 * @Date: 2021-03-11 11:46:46
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-13 14:01:03
 * @FilePath: /monkey/controllers/base.go
 */
package controllers

import (
	"html/template"
	"strconv"
	"time"

	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
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
	c.Layout = "layout/app.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navbar"] = "shared/navbar.tpl"
	c.LayoutSections["Notification"] = "shared/notification.tpl"
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
	c.Data["LayoutSetting"] = model.DefaultSetting()
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
