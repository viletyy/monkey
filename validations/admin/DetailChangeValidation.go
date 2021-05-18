/*
 * @Date: 2021-04-28 22:05:13
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-18 18:17:34
 * @FilePath: /monkey/validations/admin/DetailChangeValidation.go
 */
package admin

type DetailChangeValidation struct {
	Title     string   `form:"title" valid:"Required"`
	SubTitle  string   `form:"sub_title" valid:"Required"`
	Keywords  []string `form:"keywords"`
	Content   string   `form:"content"`
	PlateID   int      `form:"plate_id" valid:"Required"`
	Recommend bool     `form:"recommend"`
}
