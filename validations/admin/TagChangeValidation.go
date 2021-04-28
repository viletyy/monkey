/*
 * @Date: 2021-04-28 22:55:34
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 22:56:18
 * @FilePath: /monkey/validations/admin/TagChangeValidation.go
 */
package admin

type TagChangeValidation struct {
	Name string `form:"name" valid:"Required"`
}
