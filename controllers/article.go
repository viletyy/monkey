/*
 * @Date: 2021-04-22 17:41:07
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:14:57
 * @FilePath: /monkey/controllers/article.go
 */
package controllers

import (
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
)

type Article struct {
	Base
}

func (c *Article) Index() {
	articleMaps := make(map[string]interface{})
	articleMaps["detail_type"] = 0
	articleSearch := global.Search{
		Maps: articleMaps,
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 4,
		},
	}
	articleResult, articleErr := model.GetDetails(&articleSearch)
	if articleErr != nil {
		c.FlashError("加载特别推荐失败")
	}
	c.Data["Articles"] = articleResult.List

	plateMaps := make(map[string]interface{})
	plateMaps["article_recommend"] = true
	plateSearch := global.Search{
		Maps: plateMaps,
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 10,
		},
	}
	plateResult, plateErr := model.GetPlatesWithArticle(&plateSearch)
	if plateErr != nil {
		c.FlashError("加载板块失败")
	}
	c.Data["Plates"] = plateResult.List
}

func (c *Article) Show() {
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
