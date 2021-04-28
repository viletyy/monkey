/*
 * @Date: 2021-04-28 14:05:51
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 21:57:37
 * @FilePath: /monkey/model/tag.go
 */
package model

import "github.com/viletyy/monkey/global"

type Tag struct {
	global.Model
	Name string `json:"name"`
}

func GetTags(search *global.Search) (searchResult *global.SearchResult, err error) {
	var tags []Tag
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Find(&tags).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Offset(offset).Limit(limit).Find(&tags).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = tags
	return
}

func GetTagById(id int) (tag Tag, err error) {
	err = global.DB.Where("id = ?", id).First(&tag).Error
	return
}

func CreateTag(tag *Tag) (err error) {
	err = global.DB.Create(&tag).Error
	return
}

func UpdateTag(tag *Tag) (err error) {
	err = global.DB.Update(&tag).Error
	return
}

func DeleteTag(tag *Tag) (err error) {
	err = global.DB.Delete(&tag).Error
	return
}
