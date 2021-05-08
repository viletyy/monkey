/*
 * @Date: 2021-04-24 15:36:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 17:48:30
 * @FilePath: /monkey/controllers/admin/banner.go
 */
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/validations/admin"
)

type Banner struct {
	Base
}

func (c *Banner) Index() {
	maps := make(map[string]interface{})
	getName := c.GetString("name")
	if getName != "" {
		maps["name"] = getName
	}
	getPosition := c.GetString("position")
	if getPosition != "" {
		maps["position"] = getPosition
	}
	search := global.Search{
		Maps:     maps,
		PageInfo: c.PageInfo(),
	}

	if searchResult, err := model.GetBanners(&search); err == nil {
		c.ResponseWithResult(&searchResult)
	} else {
		c.FlashError("加载列表失败")
	}
}

func (c *Banner) New() {}

func (c *Banner) Create() {
	c.RedirectUrl = beego.URLFor("admin.Banner.New")
	createBanner := admin.BannerChangeValidation{}
	c.ValidatorAuto(&createBanner)

	dbBanner := model.Banner{
		Name:     createBanner.Name,
		Cover:    c.ProcessFile("cover"),
		Position: createBanner.Position,
	}

	if err := model.CreateBanner(&dbBanner); err == nil {
		c.FlashSuccess("添加Banner成功")
		c.RedirectTo("/admin/banner")
	} else {
		c.FlashError(err.Error())
		c.RedirectTo(c.RedirectUrl)
	}
}

func (c *Banner) Edit() {
	var bannerId int
	err := c.Ctx.Input.Bind(&bannerId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbBanner, err := model.GetBannerById(bannerId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["Banner"] = dbBanner
	}
}

func (c *Banner) Update() {
	var bannerId int
	err := c.Ctx.Input.Bind(&bannerId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Banner.Edit", ":id", bannerId)
	if dbBanner, err := model.GetBannerById(bannerId); err != nil {
		c.ErrorHandler(err)
	} else {
		updateBanner := admin.BannerChangeValidation{}
		c.ValidatorAuto(&updateBanner)

		//TODO 文件处理

		dbBanner.Name = updateBanner.Name
		dbBanner.Position = updateBanner.Position

		if err := model.UpdateBanner(&dbBanner); err == nil {
			c.FlashSuccess("更新Banner成功")
			c.RedirectTo("/admin/banner")
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}

func (c *Banner) Destroy() {
	var bannerId int
	err := c.Ctx.Input.Bind(&bannerId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Banner.Index")
	if dbBanner, err := model.GetBannerById(bannerId); err != nil {
		c.ErrorHandler(err)
	} else {
		if err := model.DeleteBanner(&dbBanner); err == nil {
			c.FlashSuccess("删除Banner成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}
