/*
 * @Date: 2021-04-24 15:40:34
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-26 17:53:45
 * @FilePath: /monkey/controllers/admin/user.go
 */
package admin

import (
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
)

type User struct {
	Base
}

func (c *User) Index() {
	maps := make(map[string]interface{})
	getUsername := c.GetString("username")
	if getUsername != "" {
		maps["username"] = getUsername
	}
	getIsAdmin := c.GetString("is_admin")
	if getIsAdmin != "" {
		maps["is_admin"] = getIsAdmin
	}

	search := global.Search{
		Maps:     maps,
		PageInfo: c.PageInfo(),
	}

	if searchResult, err := model.GetUsers(&search); err == nil {
		c.ResponseWithResult(searchResult)
	} else {
		c.FlashError("加载列表失败")
	}
}
