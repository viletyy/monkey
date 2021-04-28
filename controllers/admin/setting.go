/*
 * @Date: 2021-04-24 15:35:00
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 23:33:07
 * @FilePath: /monkey/controllers/admin/setting.go
 */
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/validations/admin"
)

type Setting struct {
	Base
}

func (c *Setting) Index() {
	maps := make(map[string]interface{})
	getName := c.GetString("name")
	if getName != "" {
		maps["name"] = getName
	}
	getIsCurrent := c.GetString("is_current")
	if getIsCurrent != "" {
		maps["position"] = getIsCurrent
	}
	search := global.Search{
		Maps:     maps,
		PageInfo: c.PageInfo(),
	}

	if searchResult, err := model.GetSettings(&search); err == nil {
		c.ResponseWithResult(&searchResult)
	} else {
		c.FlashError("加载列表失败")
	}
}

func (c *Setting) New() {}

func (c *Setting) Create() {
	c.RedirectUrl = beego.URLFor("admin.Setting.New")
	createSetting := admin.SettingChangeValidation{}
	c.ValidatorAuto(&createSetting)

	var keywords []model.Keyword
	var navbars model.Navbars

	//TODO keywords处理
	//TODO navbars处理

	dbSetting := model.Setting{
		Name:      createSetting.Name,
		Title:     createSetting.Title,
		Keywords:  keywords,
		Navbars:   navbars,
		Footer:    createSetting.Footer,
		IsCurrent: createSetting.IsCurrent,
	}

	if err := model.CreateSetting(&dbSetting); err == nil {
		c.FlashSuccess("添加设置成功")
		c.RedirectTo("/admin/setting")
	} else {
		c.FlashError(err.Error())
		c.RedirectTo(c.RedirectUrl)
	}
}

func (c *Setting) Edit() {
	var settingId int
	err := c.Ctx.Input.Bind(&settingId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbSetting, err := model.GetSettingById(settingId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["Setting"] = dbSetting
	}
}

func (c *Setting) Update() {
	var settingId int
	err := c.Ctx.Input.Bind(&settingId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Setting.Edit", ":id", settingId)
	if dbSetting, err := model.GetSettingById(settingId); err != nil {
		c.ErrorHandler(err)
	} else {
		updateSetting := admin.SettingChangeValidation{}
		c.ValidatorAuto(&updateSetting)

		//TODO keywords处理
		//TODO navbar处理
		var keywords []model.Keyword
		var navbars model.Navbars

		dbSetting.Name = updateSetting.Name
		dbSetting.Title = updateSetting.Title
		dbSetting.Keywords = keywords
		dbSetting.Navbars = navbars
		dbSetting.Footer = updateSetting.Footer
		dbSetting.IsCurrent = updateSetting.IsCurrent

		if err := model.UpdateSetting(&dbSetting); err == nil {
			c.FlashSuccess("更新设置成功")
			c.RedirectTo("/admin/setting")
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}

func (c *Setting) Destroy() {
	var settingId int
	err := c.Ctx.Input.Bind(&settingId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Setting.Index")
	if dbSetting, err := model.GetSettingById(settingId); err != nil {
		c.ErrorHandler(err)
	} else {
		if err := model.DeleteSetting(&dbSetting); err == nil {
			c.FlashSuccess("删除设置成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}
