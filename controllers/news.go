/*
 * @Date: 2021-04-22 17:41:35
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:15:14
 * @FilePath: /monkey/controllers/news.go
 */
package controllers

import (
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
)

type News struct {
	Base
}

func (c *News) Index() {
	newsMaps := make(map[string]interface{})
	newsMaps["detail_type"] = 1
	newsSearch := global.Search{
		Maps: newsMaps,
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 4,
		},
	}
	newsResult, newsErr := model.GetDetails(&newsSearch)
	if newsErr != nil {
		c.FlashError("加载特别推荐失败")
	}
	c.Data["News"] = newsResult.List

	plateMaps := make(map[string]interface{})
	plateMaps["news_recommend"] = true
	plateSearch := global.Search{
		Maps: plateMaps,
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 10,
		},
	}
	plateResult, plateErr := model.GetPlatesWithNews(&plateSearch)
	if plateErr != nil {
		c.FlashError("加载板块失败")
	}
	c.Data["Plates"] = plateResult.List
}

func (c *News) Show() {
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
