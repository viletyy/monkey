/*
 * @Date: 2021-04-27 20:43:43
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 22:00:00
 * @FilePath: /monkey/validations/admin/UserCreateValidation.go
 */
package admin

type UserCreateValidation struct {
	Username string `form:"username" valid:"Required;MaxSize(64)"`
	Password string `form:"password" valid:"Required;MinSize(6); MaxSize(16)"`
	IsAdmin  bool   `form:"is_admin"`
}
