/*
 * @Date: 2021-04-23 10:11:44
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 23:18:13
 * @FilePath: /monkey/controllers/admin/plate.go
 */
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/validations/admin"
)

type Plate struct {
	Base
}

func (c *Plate) Index() {
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

	if searchResult, err := model.GetPlates(&search); err == nil {
		c.ResponseWithResult(&searchResult)
	} else {
		c.FlashError("加载列表失败")
	}
}

func (c *Plate) New() {}

func (c *Plate) Create() {
	c.RedirectUrl = beego.URLFor("admin.Plate.New")
	createPlate := admin.PlateChangeValidation{}
	c.ValidatorAuto(&createPlate)

	dbPlate := model.Plate{
		Name:             createPlate.Name,
		ArticleRecommend: createPlate.ArticleRecommend,
		NewsRecommend:    createPlate.NewsRecommend,
	}

	if err := model.CreatePlate(&dbPlate); err == nil {
		c.FlashSuccess("添加板块成功")
		c.RedirectTo("/admin/plate")
	} else {
		c.FlashError(err.Error())
		c.RedirectTo(c.RedirectUrl)
	}
}

func (c *Plate) Edit() {
	var plateId int
	err := c.Ctx.Input.Bind(&plateId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbPlate, err := model.GetPlateById(plateId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["Plate"] = dbPlate
	}
}

func (c *Plate) Update() {
	var plateId int
	err := c.Ctx.Input.Bind(&plateId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Plate.Edit", ":id", plateId)
	if dbPlate, err := model.GetPlateById(plateId); err != nil {
		c.ErrorHandler(err)
	} else {
		updatePlate := admin.PlateChangeValidation{}
		c.ValidatorAuto(&updatePlate)

		dbPlate.Name = updatePlate.Name
		dbPlate.ArticleRecommend = updatePlate.ArticleRecommend
		dbPlate.NewsRecommend = updatePlate.NewsRecommend

		if err := model.UpdatePlate(&dbPlate); err == nil {
			c.FlashSuccess("更新板块成功")
			c.RedirectTo("/admin/plate")
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}

func (c *Plate) Destroy() {
	var plateId int
	err := c.Ctx.Input.Bind(&plateId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Plate.Index")
	if dbPlate, err := model.GetPlateById(plateId); err != nil {
		c.ErrorHandler(err)
	} else {
		if err := model.DeletePlate(&dbPlate); err == nil {
			c.FlashSuccess("删除板块成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}
