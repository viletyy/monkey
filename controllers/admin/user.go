/*
 * @Date: 2021-04-24 15:40:34
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 22:35:47
 * @FilePath: /monkey/controllers/admin/user.go
 */
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/validations/admin"
	"github.com/viletyy/yolk/crypt"
)

type User struct {
	Base
}

func (c *User) Index() {
	maps := make(map[string]interface{})
	getUsername := c.GetString("username")
	if getUsername != "" {
		maps["username"] = getUsername
	}
	getIsAdmin := c.GetString("is_admin")
	if getIsAdmin != "" {
		maps["is_admin"] = getIsAdmin
	}

	search := global.Search{
		Maps:     maps,
		PageInfo: c.PageInfo(),
	}

	if searchResult, err := model.GetUsers(&search); err == nil {
		c.ResponseWithResult(&searchResult)
	} else {
		c.FlashError("加载列表失败")
	}
}

func (c *User) New() {

}

func (c *User) Create() {
	c.RedirectUrl = beego.URLFor("admin.User.New")
	createUser := admin.UserCreateValidation{}
	c.ValidatorAuto(&createUser)

	if _, isExist := model.ExistByUsername(createUser.Username); isExist {
		c.FlashError("该用户名已被使用")
		c.RedirectTo(c.RedirectUrl)
	}

	dbUser := model.User{
		Username: createUser.Username,
		Password: crypt.Md5Encode(createUser.Password),
		IsAdmin:  createUser.IsAdmin,
	}

	if err := model.CreateUser(&dbUser); err == nil {
		c.FlashSuccess("添加用户成功")
		c.RedirectTo("/admin/user")
	} else {
		c.FlashError(err.Error())
		c.RedirectTo(c.RedirectUrl)
	}
}

func (c *User) Edit() {
	var userId int
	err := c.Ctx.Input.Bind(&userId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbUser, err := model.GetUserById(userId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["User"] = dbUser
	}
}

func (c *User) Update() {
	var userId int
	err := c.Ctx.Input.Bind(&userId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.User.Edit", ":id", userId)
	if dbUser, err := model.GetUserById(userId); err != nil {
		c.ErrorHandler(err)
	} else {
		updateUser := admin.UserUpdateValidation{}
		c.ValidatorAuto(&updateUser)

		dbUser.Username = updateUser.Username
		dbUser.IsAdmin = updateUser.IsAdmin

		if err := model.UpdateUser(&dbUser); err == nil {
			c.FlashSuccess("更新成功")
			c.RedirectTo("/admin/user")
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}

func (c *User) Reset() {
	var userId int
	_ = c.Ctx.Input.Bind(&userId, ":id")
	c.RedirectUrl = beego.URLFor("admin.User.Index")
	if dbUser, err := model.GetUserById(userId); err != nil {
		c.ErrorHandler(err)
	} else {
		dbUser.Password = crypt.Md5Encode("123456")
		if err := model.UpdateUser(&dbUser); err == nil {
			c.FlashSuccess("重置密码成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}

}

func (c *User) Destroy() {
	var userId int
	_ = c.Ctx.Input.Bind(&userId, ":id")
	c.RedirectUrl = beego.URLFor("admin.User.Index")
	if dbUser, err := model.GetUserById(userId); err != nil {
		c.ErrorHandler(err)
	} else {
		if err := model.DeleteUser(&dbUser); err == nil {
			c.FlashSuccess("删除用户成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}
