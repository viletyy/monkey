/*
 * @Date: 2021-03-12 16:28:24
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 16:31:50
 * @FilePath: /monkey/utils/auth.go
 */
package utils

import (
	"strconv"

	"github.com/viletyy/yolk/crypt"

	beego "github.com/beego/beego/v2/server/web"
)

const (
	AuthSecret = "golangmore"
)

func AuthSign(id int, username, password string) string {
	s := strconv.Itoa(id) + username + beego.Substr(password, 0, 13)
	return crypt.Sha256Encode(s, AuthSecret)
}

func AuthSignCheck(id int, username, password, sign string) bool {
	s := AuthSign(id, username, password)
	return s == sign
}
