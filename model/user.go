/*
 * @Date: 2021-03-09 09:57:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 21:32:53
 * @FilePath: /monkey/model/user.go
 */
package model

import (
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/yolk/crypt"
)

type User struct {
	global.Model
	Username string   `json:"username"`
	Password string   `json:"-"`
	IsAdmin  bool     `json:"is_admin" gorm:"default: false"`
	Details  []Detail `json:"details"`
	Files    []File   `json:"files"`
}

func GetUsers(search *global.Search) (searchResult global.SearchResult, err error) {
	var users []User
	offset := search.PageSize * (search.Page - 1)
	limit := search.PageSize
	err = global.DB.Where(search.Maps).Find(&users).Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = global.DB.Where(search.Maps).Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return
	}
	searchResult.Page = search.PageInfo.Page
	searchResult.PageSize = search.PageInfo.PageSize
	searchResult.List = users
	return
}

func GetUserById(id int) (user User, err error) {
	err = global.DB.Where("id = ?", id).First(&user).Error
	return
}

func ExistByUsername(username string) (user User, isExist bool) {
	global.DB.Where("username = ?", username).First(&user)

	return user, user.ID > 0
}

func CreateUser(user *User) (err error) {
	err = global.DB.Create(&user).Error

	return
}

func UpdateUser(user *User) (err error) {
	err = global.DB.Save(&user).Error
	return
}

func DeleteUser(user *User) (err error) {
	err = global.DB.Delete(&user).Error
	return
}

func (user *User) CheckPassword(password string) bool {
	return crypt.Md5Check(password, user.Password)
}
