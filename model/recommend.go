/*
 * @Date: 2021-04-28 16:27:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 21:37:38
 * @FilePath: /monkey/model/recommend.go
 */
package model

import "github.com/viletyy/monkey/global"

type Recommend struct {
	global.Model
	Name string `json:"name"`
	Link string `json:"link"`
}

func GetRecommends(search *global.Search) (searchResult global.SearchResult, err error) {
	var recommends []Recommend
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Find(&recommends).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Offset(offset).Limit(limit).Find(&recommends).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = recommends
	return
}

func GetRecommendById(id int) (recommend Recommend, err error) {
	err = global.DB.Where("id = ?", id).First(&recommend).Error
	return
}

func CreateRecommend(recommend *Recommend) (err error) {
	err = global.DB.Create(&recommend).Error
	return
}

func UpdateRecommend(recommend *Recommend) (err error) {
	err = global.DB.Save(&recommend).Error
	return
}

func DeleteRecommend(recommend *Recommend) (err error) {
	err = global.DB.Delete(&recommend).Error
	return
}
