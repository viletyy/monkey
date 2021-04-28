/*
 * @Date: 2021-04-24 15:36:51
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 23:20:34
 * @FilePath: /monkey/controllers/admin/recommend.go
 */
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/validations/admin"
)

type Recommend struct {
	Base
}

func (c *Recommend) Index() {
	maps := make(map[string]interface{})
	getName := c.GetString("name")
	if getName != "" {
		maps["name"] = getName
	}
	search := global.Search{
		Maps:     maps,
		PageInfo: c.PageInfo(),
	}

	if searchResult, err := model.GetRecommends(&search); err == nil {
		c.ResponseWithResult(&searchResult)
	} else {
		c.FlashError("加载列表失败")
	}
}

func (c *Recommend) New() {}

func (c *Recommend) Create() {
	c.RedirectUrl = beego.URLFor("admin.Recommend.New")
	createRecommend := admin.RecommendChangeValidation{}
	c.ValidatorAuto(&createRecommend)

	dbRecommend := model.Recommend{
		Name: createRecommend.Name,
		Link: createRecommend.Link,
	}

	if err := model.CreateRecommend(&dbRecommend); err == nil {
		c.FlashSuccess("添加推荐成功")
		c.RedirectTo("/admin/recommend")
	} else {
		c.FlashError(err.Error())
		c.RedirectTo(c.RedirectUrl)
	}
}

func (c *Recommend) Edit() {
	var recommendId int
	err := c.Ctx.Input.Bind(&recommendId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbRecommend, err := model.GetRecommendById(recommendId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["Recommend"] = dbRecommend
	}
}

func (c *Recommend) Update() {
	var recommendId int
	err := c.Ctx.Input.Bind(&recommendId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Recommend.Edit", ":id", recommendId)
	if dbRecommend, err := model.GetRecommendById(recommendId); err != nil {
		c.ErrorHandler(err)
	} else {
		updateRecommend := admin.RecommendChangeValidation{}
		c.ValidatorAuto(&updateRecommend)

		dbRecommend.Name = updateRecommend.Name
		dbRecommend.Link = updateRecommend.Link

		if err := model.UpdateRecommend(&dbRecommend); err == nil {
			c.FlashSuccess("更新推荐成功")
			c.RedirectTo("/admin/recommend")
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}

func (c *Recommend) Destroy() {
	var recommendId int
	err := c.Ctx.Input.Bind(&recommendId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Recommend.Index")
	if dbRecommend, err := model.GetRecommendById(recommendId); err != nil {
		c.ErrorHandler(err)
	} else {
		if err := model.DeleteRecommend(&dbRecommend); err == nil {
			c.FlashSuccess("删除推荐成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}
