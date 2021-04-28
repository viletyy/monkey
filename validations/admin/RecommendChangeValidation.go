/*
 * @Date: 2021-04-28 22:48:00
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 22:50:24
 * @FilePath: /monkey/validations/admin/RecommendChangeValidation.go
 */
package admin

type RecommendChangeValidation struct {
	Name string `form:"name" valid:"Required"`
	Link string `form:"link" valid:"Required"`
}
