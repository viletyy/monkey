/*
 * @Date: 2021-03-12 17:25:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-26 13:53:53
 * @FilePath: /monkey/validations/UserRegisterValidation.go
 */
package validations

import (
	"github.com/viletyy/monkey/model"

	"github.com/beego/beego/v2/core/validation"
)

type UserRegisterValidation struct {
	Username             string `form:"username" valid:"Required; MaxSize(64)"`
	Password             string `form:"password" valid:"Required; MinSize(6); MaxSize(16)"`
	PasswordConfirmation string `form:"password_confirmation"`
}

func (urv *UserRegisterValidation) Valid(valid *validation.Validation) {
	if urv.PasswordConfirmation != "" && urv.Password != urv.PasswordConfirmation {
		_ = valid.SetError("password", "两次输入密码不一致")
		return
	}

	if _, isExsit := model.ExistByUsername(urv.Username); !isExsit {
		_ = valid.SetError("username", "用户名已经存在")
		return
	}
}
