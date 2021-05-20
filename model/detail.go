/*
 * @Date: 2021-04-28 14:08:12
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:48:55
 * @FilePath: /monkey/model/detail.go
 */
package model

import "github.com/viletyy/monkey/global"

type DetailType int

const (
	DetailArticle DetailType = iota
	DetailNews
)

type Detail struct {
	global.Model
	Title      string          `json:"title"`
	SubTitle   string          `json:"sub_title"`
	Keywords   global.Keywords `json:"keywords" sql:"TYPE:json"`
	Cover      File            `json:"cover" gorm:"polymorphic:Owner"`
	Content    string          `json:"content"`
	DetailType DetailType      `json:"detail_type"`
	Tags       []Tag           `json:"tags" gorm:"many2many:detail_tags"`
	UserID     uint            `json:"user_id"`
	User       User            `json:"user" gorm:"association_autoupdate:false"`
	PlateID    uint            `json:"plate_id"`
	Plate      Plate           `json:"plate" gorm:"association_autoupdate:false"`
	Recommend  bool            `json:"recommend"`
}

func GetDetails(search *global.Search) (searchResult global.SearchResult, err error) {
	var details []Detail
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize

	err = global.DB.Where(search.Maps).Find(&details).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Preload("Tags").Preload("Plate").Preload("Cover").Preload("User").Offset(offset).Limit(limit).Find(&details).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = details
	return
}

func GetDetailTotalCount(search *global.Search) (total int64, err error) {
	err = global.DB.Model(&Detail{}).Where(search.Maps).Count(&total).Error
	return
}

func GetDetailById(id int) (detail Detail, err error) {
	err = global.DB.Where("id = ?", id).Preload("Tags").Preload("Plate").Preload("Cover").Preload("User").First(&detail).Error
	return
}

func CreateDetail(detail *Detail) (err error) {
	err = global.DB.Create(&detail).Error
	return
}

func UpdateDetail(detail *Detail) (err error) {
	err = global.DB.Save(&detail).Error
	return
}

func DeleteDetail(detail *Detail) (err error) {
	err = global.DB.Delete(&detail).Error
	return
}
