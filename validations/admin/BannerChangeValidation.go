/*
 * @Date: 2021-04-28 22:40:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 22:49:44
 * @FilePath: /monkey/validations/admin/BannerChangeValidation.go
 */
package admin

import "mime/multipart"

type BannerChangeValidation struct {
	Name     string         `form:"name" valid:"Required" `
	Cover    multipart.File `form:"cover" valid:"Required"`
	Position int            `form:"position" valid:"Required"`
}
