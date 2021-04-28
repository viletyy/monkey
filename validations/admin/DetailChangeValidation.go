/*
 * @Date: 2021-04-28 22:05:13
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 22:30:14
 * @FilePath: /monkey/validations/admin/DetailChangeValidation.go
 */
package admin

import "mime/multipart"

type DetailChangeValidation struct {
	Title     string         `form:"title" valid:"Required"`
	SubTitle  string         `form:"sub_title" valid:"Required"`
	Cover     multipart.File `form:"cover" valid:"Required"`
	Content   string         `form:"content"`
	Tags      []interface{}  `form:"tags"`
	PlateID   int            `form:"plate_id" valid:"Required"`
	Recommend bool           `form:"recommend"`
}
