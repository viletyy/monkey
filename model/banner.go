/*
 * @Date: 2021-04-28 16:28:12
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-11 14:23:37
 * @FilePath: /monkey/model/banner.go
 */
package model

import (
	"github.com/viletyy/monkey/global"
)

type Banner struct {
	global.Model
	Name     string `json:"name"`
	Cover    File   `json:"cover" gorm:"polymorphic:Owner"`
	Position int    `json:"position"`
}

func GetBanners(search *global.Search) (searchResult global.SearchResult, err error) {
	var banners []Banner
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Find(&banners).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Preload("Cover").Offset(offset).Limit(limit).Find(&banners).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = banners
	return
}

func GetBannerById(id int) (banner Banner, err error) {
	err = global.DB.Where("id = ?", id).Preload("Cover").First(&banner).Error
	return
}

func CreateBanner(banner *Banner) (err error) {
	err = global.DB.Create(&banner).Error
	return
}

func UpdateBanner(banner *Banner) (err error) {
	err = global.DB.Save(&banner).Error
	return
}

func DeleteBanner(banner *Banner) (err error) {
	err = global.DB.Delete(&banner).Error
	return
}
