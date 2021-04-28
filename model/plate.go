/*
 * @Date: 2021-04-28 15:26:46
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 21:37:16
 * @FilePath: /monkey/model/plate.go
 */
package model

import "github.com/viletyy/monkey/global"

type Plate struct {
	global.Model
	Name         string   `json:"name"`
	Recommend    bool     `json:"recommend"`
	ArticleCount int      `json:"article_count" gorm:"default: 0"`
	NewsCount    int      `json:"news_count" gorm:"default: 0"`
	Details      []Detail `json:"details"`
}

func GetPlates(search *global.Search) (searchResult global.SearchResult, err error) {
	var plates []Plate
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Find(&plates).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Offset(offset).Limit(limit).Find(&plates).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = plates
	return
}

func GetPlateById(id int) (plate Plate, err error) {
	err = global.DB.Where("id = ?", id).First(&plate).Error
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
