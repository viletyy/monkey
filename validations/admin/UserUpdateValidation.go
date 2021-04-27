/*
 * @Date: 2021-04-27 21:58:59
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 23:41:42
 * @FilePath: /monkey/validations/admin/UserUpdateValidation.go
 */
package admin

type UserUpdateValidation struct {
	Username string `form:"username" valid:"Required;MaxSize(64)"`
	IsAdmin  bool   `form:"is_admin" `
}
