/*
 * @Date: 2021-04-28 16:40:43
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 21:36:57
 * @FilePath: /monkey/model/file.go
 */
package model

import "github.com/viletyy/monkey/global"

type File struct {
	global.Model
	OwnerID     int    `json:"owner_id"`
	OwnerType   int    `json:"owner_type"`
	UserID      int    `json:"user_id"`
	User        User   `json:"user" gorm:"association_autoupdate:false"`
	Name        string `json:"name"`
	DiskName    string `json:"disk_name"`
	SavePath    string `json:"save_path"`
	Size        int    `json:"size"`
	ContentType string `json:"content_type"`
}

func GetFiles(search *global.Search) (searchResult global.SearchResult, err error) {
	var files []File
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Find(&files).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Offset(offset).Limit(limit).Find(&files).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = files
	return
}

func GetFileById(id int) (file File, err error) {
	err = global.DB.Where("id = ?", id).Error
	return
}

func CreateFile(file *File) (err error) {
	err = global.DB.Create(&file).Error
	return
}

func UpdateFile(file *File) (err error) {
	err = global.DB.Save(&file).Error
	return
}

func DeleteFile(file *File) (err error) {
	err = global.DB.Delete(&file).Error
	return
}
