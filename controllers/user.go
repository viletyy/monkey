/*
 * @Date: 2021-03-11 10:26:08
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 16:49:51
 * @FilePath: /monkey/controllers/user.go
 */
package controllers

import (
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/utils"
	"github.com/viletyy/monkey/validations"
	"github.com/viletyy/yolk"
	"github.com/viletyy/yolk/crypt"

	beego "github.com/beego/beego/v2/server/web"
)

type User struct {
	Base
}

func (c *User) Login() {
	if c.CurrentLoginUser != nil {
		c.RedirectTo("/")
	}
}

func (c *User) LoginHandle() {
	c.RedirectUrl = beego.URLFor("User.Login")
	loginUser := validations.UserLoginValidation{}
	c.ValidatorAuto(&loginUser)

	dbUser, isExist := model.ExistByUsername(loginUser.Username)
	if !isExist {
		c.FlashError("用户不存在")
		c.RedirectTo(c.RedirectUrl)
	}

	if !dbUser.CheckPassword(loginUser.Password) {
		c.FlashError("密码不正确")
		c.RedirectTo(c.RedirectUrl)
	}

	if c.GetString("remember_me") != "" {
		c.Ctx.SetCookie("login_user_id", yolk.ToString(dbUser.ID), 3600*24*7)
		c.Ctx.SetCookie("login_user_sign", utils.AuthSign(int(dbUser.ID), dbUser.Username, dbUser.Password), 3600*24*7)
	} else {
		c.Ctx.SetCookie("login_user_id", yolk.ToString(dbUser.ID), 3600)
		c.Ctx.SetCookie("login_user_sign", utils.AuthSign(int(dbUser.ID), dbUser.Username, dbUser.Password), 3600)
	}

	c.FlashSuccess("登陆成功")
	c.RedirectTo("/")
}

func (c *User) Register() {

}

func (c *User) RegisterHandle() {
	c.RedirectUrl = beego.URLFor("User.Register")
	registerUser := validations.UserRegisterValidation{}
	c.ValidatorAuto(&registerUser)

	dbUser := model.User{
		Username: registerUser.Username,
		Password: crypt.Md5Encode(registerUser.Password),
	}

	if err := model.CreateUser(dbUser); err != nil {
		c.FlashError("注册失败")
		c.RedirectTo(c.RedirectUrl)
	} else {
		c.FlashSuccess("注册成功")
		c.RedirectTo(beego.URLFor("User.Login"))
	}
}

func (c *User) Loginout() {
	c.CurrentLoginUser = nil
	c.Ctx.SetCookie("login_user_id", "", -3600)
	c.Ctx.SetCookie("login_user_sign", "", -3600)
	c.FlashSuccess("已安全登出")
	c.RedirectTo(c.RedirectUrl)
}
