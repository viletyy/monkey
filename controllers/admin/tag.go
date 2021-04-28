/*
 * @Date: 2021-04-23 10:14:08
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 23:34:49
 * @FilePath: /monkey/controllers/admin/tag.go
 */
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/validations/admin"
)

type Tag struct {
	Base
}

func (c *Tag) Index() {
	maps := make(map[string]interface{})
	getName := c.GetString("name")
	if getName != "" {
		maps["name"] = getName
	}
	search := global.Search{
		Maps:     maps,
		PageInfo: c.PageInfo(),
	}

	if searchResult, err := model.GetTags(&search); err == nil {
		c.ResponseWithResult(&searchResult)
	} else {
		c.FlashError("加载列表失败")
	}
}

func (c *Tag) New() {}

func (c *Tag) Create() {
	c.RedirectUrl = beego.URLFor("admin.Tag.New")
	createTag := admin.TagChangeValidation{}
	c.ValidatorAuto(&createTag)

	dbTag := model.Tag{
		Name: createTag.Name,
	}

	if err := model.CreateTag(&dbTag); err == nil {
		c.FlashSuccess("添加标签成功")
		c.RedirectTo("/admin/tag")
	} else {
		c.FlashError(err.Error())
		c.RedirectTo(c.RedirectUrl)
	}
}

func (c *Tag) Edit() {
	var tagId int
	err := c.Ctx.Input.Bind(&tagId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbTag, err := model.GetTagById(tagId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["Tag"] = dbTag
	}
}

func (c *Tag) Update() {
	var tagId int
	err := c.Ctx.Input.Bind(&tagId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Tag.Edit", ":id", tagId)
	if dbTag, err := model.GetTagById(tagId); err != nil {
		c.ErrorHandler(err)
	} else {
		updateTag := admin.TagChangeValidation{}
		c.ValidatorAuto(&updateTag)

		dbTag.Name = updateTag.Name

		if err := model.UpdateTag(&dbTag); err == nil {
			c.FlashSuccess("更新标签成功")
			c.RedirectTo("/admin/tag")
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}

func (c *Tag) Destroy() {
	var tagId int
	err := c.Ctx.Input.Bind(&tagId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Tag.Index")
	if dbTag, err := model.GetTagById(tagId); err != nil {
		c.ErrorHandler(err)
	} else {
		if err := model.DeleteTag(&dbTag); err == nil {
			c.FlashSuccess("删除标签成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}
