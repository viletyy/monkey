/*
 * @Date: 2021-04-28 22:40:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 17:13:59
 * @FilePath: /monkey/validations/admin/BannerChangeValidation.go
 */
package admin

type BannerChangeValidation struct {
	Name     string `form:"name" valid:"Required" `
	Position int    `form:"position" valid:"Required"`
}
