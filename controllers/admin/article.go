/*
 * @Date: 2021-04-23 10:12:31
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 23:05:10
 * @FilePath: /monkey/controllers/admin/article.go
 */
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
	"github.com/viletyy/monkey/validations/admin"
)

type Article struct {
	Base
}

func (c *Article) Index() {
	maps := make(map[string]interface{})
	maps["detail_type"] = model.DetailArticle
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

func (c *Article) New() {

}

func (c *Article) Create() {
	c.RedirectUrl = beego.URLFor("admin.Article.New")
	createArticle := admin.DetailChangeValidation{}
	c.ValidatorAuto(&createArticle)

	//TODO文件处理
	//tags处理
	var tags []model.Tag

	dbArticle := model.Detail{
		Title:      createArticle.Title,
		SubTitle:   createArticle.SubTitle,
		Content:    createArticle.Content,
		DetailType: model.DetailArticle,
		Tags:       tags,
		UserID:     c.CurrentLoginUser.ID,
		PlateID:    uint(createArticle.PlateID),
		Recommend:  createArticle.Recommend,
	}

	if err := model.CreateDetail(&dbArticle); err == nil {
		c.FlashSuccess("添加文章成功")
		c.RedirectTo("/admin/article")
	} else {
		c.FlashError(err.Error())
		c.RedirectTo(c.RedirectUrl)
	}
}

func (c *Article) Edit() {
	var articleId int
	err := c.Ctx.Input.Bind(&articleId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbArticle, err := model.GetDetailById(articleId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["Article"] = dbArticle
	}
}

func (c *Article) Update() {
	var articleId int
	err := c.Ctx.Input.Bind(&articleId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Article.Edit", ":id", articleId)
	if dbArticle, err := model.GetDetailById(articleId); err != nil {
		c.ErrorHandler(err)
	} else {
		updateArticle := admin.DetailChangeValidation{}
		c.ValidatorAuto(&updateArticle)

		//TODO 文件处理
		//TODO tags处理

		var tags []model.Tag

		dbArticle.Title = updateArticle.Title
		dbArticle.SubTitle = updateArticle.SubTitle
		dbArticle.Content = updateArticle.Content
		dbArticle.Tags = tags
		dbArticle.PlateID = uint(updateArticle.PlateID)
		dbArticle.Recommend = updateArticle.Recommend

		if err := model.UpdateDetail(&dbArticle); err == nil {
			c.FlashSuccess("更新成功")
			c.RedirectTo("/admin/article")
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}

func (c *Article) Destroy() {
	var articleId int
	err := c.Ctx.Input.Bind(&articleId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	c.RedirectUrl = beego.URLFor("admin.Article.Index")
	if dbArticle, err := model.GetDetailById(articleId); err != nil {
		c.ErrorHandler(err)
	} else {
		if err := model.DeleteDetail(&dbArticle); err == nil {
			c.FlashSuccess("删除文章成功")
			c.RedirectTo(c.RedirectUrl)
		} else {
			c.FlashError(err.Error())
			c.RedirectTo(c.RedirectUrl)
		}
	}
}
