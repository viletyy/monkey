/*
 * @Date: 2021-04-23 10:13:24
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 16:22:15
 * @FilePath: /monkey/controllers/admin/news.go
 */
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/validations/admin"
)

type News struct {
	Base
}

func (c *News) Index() {
	maps := make(map[string]interface{})
	maps["detail_type"] = model.DetailNews
	getTitle := c.GetString("title")
	if getTitle != "" {
		maps["title"] = getTitle
	}
	getUserId := c.GetString("user_id")
	if getUserId != "" {
		maps["user_id"] = getUserId
	}
	getPlateId := c.GetString("plate_id")
	if getPlateId != "" {
		maps["plate_id"] = getPlateId
	}
	getRecommend := c.GetString("recommend")
	if getRecommend != "" {
		maps["recommend"] = getRecommend
	}
	//TODO标签搜索
	search := global.Search{
		Maps:     maps,
		PageInfo: c.PageInfo(),
	}

	if searchResult, err := model.GetDetails(&search); err == nil {
		c.ResponseWithResult(&searchResult)
	} else {
		c.FlashError("加载列表失败")
	}
}

func (c *News) New() {
	c.getPlates()
}

func (c *News) Create() {
	c.RedirectUrl = beego.URLFor("admin.News.New")
	createNews := admin.DetailChangeValidation{}
	c.ValidatorAuto(&createNews)

	//tags处理
	var tags []model.Tag

	dbNews := model.Detail{
		Title:      createNews.Title,
		SubTitle:   createNews.SubTitle,
		Keywords:   createNews.Keywords,
		Cover:      c.ProcessFile("cover"),
		Content:    createNews.Content,
		DetailType: model.DetailNews,
		Tags:       tags,
		UserID:     c.CurrentLoginUser.ID,
		PlateID:    uint(createNews.PlateID),
		Recommend:  createNews.Recommend,
	}

	if err := model.CreateDetail(&dbNews); err == nil {
		c.FlashSuccess("添加文章成功")
		c.RedirectTo("/admin/news")
	} else {
		c.FlashError(err.Error())
		c.RedirectTo(c.RedirectUrl)
	}
}

func (c *News) Edit() {
	c.getPlates()
	var newsId int
	err := c.Ctx.Input.Bind(&newsId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbNews, err := model.GetDetailById(newsId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["News"] = dbNews
	}
}

func (c *News) Update() {
	var newsId int
	err := c.Ctx.Input.Bind(&newsId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.News.Edit", ":id", newsId)
	if dbNews, err := model.GetDetailById(newsId); err != nil {
		c.ErrorHandler(err)
	} else {
		updateNews := admin.DetailChangeValidation{}
		c.ValidatorAuto(&updateNews)

		if dbPlate, err := model.GetPlateById(updateNews.PlateID); err != nil {
			c.ErrorHandler(err)
		} else {
			//TODO tags处理

			var tags []model.Tag

			file := c.ProcessFile("cover")
			if file.Size != 0 {
				dbNews.Cover = file
			}

			dbNews.Title = updateNews.Title
			dbNews.SubTitle = updateNews.SubTitle
			dbNews.Content = updateNews.Content
			dbNews.Tags = tags
			dbNews.Plate = dbPlate
			dbNews.Recommend = updateNews.Recommend

			if err := model.UpdateDetail(&dbNews); err == nil {
				c.FlashSuccess("更新成功")
				c.RedirectTo("/admin/news")
			} else {
				c.FlashError(err.Error())
				c.RedirectTo(c.RedirectUrl)
			}
		}
	}
}

func (c *News) Destroy() {
	var newsId int
	err := c.Ctx.Input.Bind(&newsId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.News.Index")
	if dbNews, err := model.GetDetailById(newsId); err != nil {
		c.ErrorHandler(err)
	} else {
		if err := model.DeleteDetail(&dbNews); err == nil {
			c.FlashSuccess("删除文章成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}

func (c *News) getPlates() {
	search := global.Search{
		Maps: make(map[string]interface{}),
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 200,
		},
	}
	if searchResult, err := model.GetPlates(&search); err == nil {
		c.Data["Plates"] = searchResult.List
	} else {
		c.FlashError("加载板块失败")
	}
}
