/*
 * @Date: 2021-04-28 16:01:21
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-13 14:14:11
 * @FilePath: /monkey/model/setting.go
 */
package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/viletyy/monkey/global"
)

type Navbar struct {
	Name   string `json:"name"`
	Link   string `json:"link"`
	IsShow bool   `json:"is_show"`
}

type Navbars []Navbar

func (n Navbars) Value() (driver.Value, error) {
	b, err := json.Marshal(n)
	return string(b), err
}

func (n *Navbars) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), n)
}

type Setting struct {
	global.Model
	Name        string          `json:"name"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Keywords    global.Keywords `json:"keywords" sql:"TYPE:json"`
	Navbars     Navbars         `json:"navbars" sql:"TYPE:json"`
	Footer      string          `json:"footer"`
	IsCurrent   bool            `json:"is_current"`
}

func DefaultSetting() *Setting {
	var currentSetting Setting
	global.DB.Where("is_current = ?", true).First(&currentSetting)
	if currentSetting.ID > 0 {
		return &currentSetting
	} else {
		defaultHttpport, defaultHttpportErr := beego.AppConfig.Int("httpport")
		if defaultHttpportErr != nil {
			panic("Get Config Error: httpport")
		}
		return &Setting{
			Name:        fmt.Sprintf("localhost:%d", defaultHttpport),
			Title:       "Monkey",
			Keywords:    []string{"Golang", "Beego", "Gorm"},
			Description: "奋力向前的程序🐶的个人博客，在平时开发中遇到的技术难关以及推荐的资讯",
			Navbars: []Navbar{
				{Name: "首页", Link: "/", IsShow: true},
				{Name: "板块", Link: "/plate", IsShow: true},
				{Name: "资讯", Link: "/news", IsShow: true},
				{Name: "文章", Link: "/article", IsShow: true},
			},
			IsCurrent: true,
		}
	}
}

func GetSettings(search *global.Search) (searchResult global.SearchResult, err error) {
	var settings []Setting
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Find(&settings).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Offset(offset).Limit(limit).Find(&settings).Error
	if err != nil {
		return
	}
	searchResult.Page = search.Page
	searchResult.PageSize = search.PageSize
	searchResult.List = settings
	return
}

func GetSettingById(id int) (setting Setting, err error) {
	err = global.DB.Where("id = ?", id).First(&setting).Error
	return
}

func CreateSetting(setting *Setting) (err error) {
	tx := global.DB.Begin()
	err = tx.Create(&setting).Error
	tx.Model(&Setting{}).Where("is_current = ? AND id != ?", true, setting.ID).Update("is_current", false)
	tx.Commit()
	return
}

func UpdateSetting(setting *Setting) (err error) {
	tx := global.DB.Begin()
	err = tx.Save(&setting).Error
	tx.Model(&Setting{}).Where("is_current = ? AND id != ?", true, setting.ID).Update("is_current", false)
	tx.Commit()
	return
}

func DeleteSetting(setting *Setting) (err error) {
	err = global.DB.Delete(&setting).Error
	return
}
