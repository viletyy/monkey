/*
 * @Date: 2021-04-28 15:26:46
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:12:59
 * @FilePath: /monkey/model/plate.go
 */
package model

import "github.com/viletyy/monkey/global"

type Plate struct {
	global.Model
	Name             string   `json:"name"`
	ArticleRecommend bool     `json:"article_recommend" gorm:"default: false"`
	ArticleCount     int      `json:"article_count" gorm:"default: 0"`
	NewsRecommend    bool     `json:"news_recommend" gorm:"default: false"`
	NewsCount        int      `json:"news_count" gorm:"default: 0"`
	Details          []Detail `json:"details"`
}

func GetPlates(search *global.Search) (searchResult global.SearchResult, err error) {
	var plates []Plate
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Preload("Details").Find(&plates).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Preload("Details").Offset(offset).Limit(limit).Find(&plates).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = plates
	return
}

func GetPlatesWithArticle(search *global.Search) (searchResult global.SearchResult, err error) {
	var plates []Plate
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Preload("Details", "detail_type = ?", DetailArticle).Find(&plates).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Preload("Details", "detail_type = ?", DetailArticle).Offset(offset).Limit(limit).Find(&plates).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = plates
	return
}

func GetPlatesWithNews(search *global.Search) (searchResult global.SearchResult, err error) {
	var plates []Plate
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Preload("Details", "detail_type = ?", DetailNews).Find(&plates).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Preload("Details", "detail_type = ?", DetailNews).Offset(offset).Limit(limit).Find(&plates).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = plates
	return
}

func GetPlateById(id int) (plate Plate, err error) {
	err = global.DB.Where("id = ?", id).Preload("Details").First(&plate).Error
	return
}

func CreatePlate(plate *Plate) (err error) {
	err = global.DB.Create(&plate).Error
	return
}

func UpdatePlate(plate *Plate) (err error) {
	err = global.DB.Save(&plate).Error
	return
}

func DeletePlate(plate *Plate) (err error) {
	err = global.DB.Delete(&plate).Error
	return
}
