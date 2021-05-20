/*
 * @Date: 2021-03-11 16:33:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:34:04
 * @FilePath: /monkey/controllers/index.go
 */
package controllers

import (
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
)

type Index struct {
	Base
}

func (c *Index) Index() {
	bannerSearch := global.Search{
		Maps: make(map[string]interface{}),
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 5,
		},
	}
	bannerResult, bannerErr := model.GetBanners(&bannerSearch)
	if bannerErr != nil {
		c.FlashError("加载Banner失败")
	}
	c.Data["Banners"] = bannerResult.List

	recommendSearch := global.Search{
		Maps: make(map[string]interface{}),
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 6,
		},
	}
	recommendResult, recommendErr := model.GetRecommends(&recommendSearch)
	if recommendErr != nil {
		c.FlashError("加载推荐资源失败")
	}
	c.Data["Recommends"] = recommendResult.List

	detailSearch := global.Search{
		Maps: make(map[string]interface{}),
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 4,
		},
	}
	detailResult, detailErr := model.GetDetails(&detailSearch)
	if detailErr != nil {
		c.FlashError("加载特别推荐失败")
	}
	c.Data["Details"] = detailResult.List

	plateMaps := make(map[string]interface{})
	plateMaps["article_recommend"] = true
	plateMaps["news_recommend"] = true
	plateSearch := global.Search{
		Maps: plateMaps,
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 4,
		},
	}
	plateResult, plateErr := model.GetPlates(&plateSearch)
	if plateErr != nil {
		c.FlashError("加载板块失败")
	}
	c.Data["Plates"] = plateResult.List

	totalMaps := make(map[string]interface{})
	totalMaps["detail_type"] = model.DetailArticle
	totalSearch := global.Search{
		Maps: totalMaps,
	}
	articleTotalCount, articleTotalErr := model.GetDetailTotalCount(&totalSearch)
	if articleTotalErr != nil {
		c.FlashError("加载文章数失败")
	}
	c.Data["ArticleTotalCount"] = articleTotalCount
	totalMaps["detail_type"] = model.DetailNews
	newsTotalCount, newsTotalErr := model.GetDetailTotalCount(&totalSearch)
	if newsTotalErr != nil {
		c.FlashError("加载资讯数失败")
	}
	c.Data["NewsTotalCount"] = newsTotalCount
}

func (c *Index) Search() {
	maps := make(map[string]interface{})
	// TODO 关键词搜索
	search := global.Search{
		Maps:     maps,
		PageInfo: c.PageInfoBox(),
	}

	detailResult, detailErr := model.GetDetails(&search)
	if detailErr != nil {
		c.FlashError("加载列表失败")
	}

	c.ResponseWithResult(&detailResult)
}
