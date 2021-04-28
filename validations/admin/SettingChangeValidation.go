/*
 * @Date: 2021-04-28 22:49:00
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 22:54:08
 * @FilePath: /monkey/validations/admin/SettingChangeValidation.go
 */
package admin

type SettingChangeValidation struct {
	Name      string        `form:"name" valid:"Required"`
	Title     string        `form:"title" valid:"Required"`
	Keywords  []string      `form:"keywords"`
	Navbars   []interface{} `form:"navbars"`
	Footer    string        `form:"footer"`
	IsCurrent bool          `form:"is_current"`
}
