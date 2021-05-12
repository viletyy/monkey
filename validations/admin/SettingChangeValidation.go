/*
 * @Date: 2021-04-28 22:49:00
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-11 18:45:30
 * @FilePath: /monkey/validations/admin/SettingChangeValidation.go
 */
package admin

import (
	"github.com/viletyy/monkey/model"
)

type SettingChangeValidation struct {
	Name        string        `form:"name" valid:"Required"`
	Title       string        `form:"title" valid:"Required"`
	Description string        `form:"description"`
	Keywords    []string      `form:"keywords"`
	Navbars     model.Navbars `form:"navbars"`
	Footer      string        `form:"footer"`
	IsCurrent   bool          `form:"is_current"`
}
