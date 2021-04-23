/*
 * @Date: 2021-04-22 17:41:07
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-23 13:58:49
 * @FilePath: /monkey/controllers/article.go
 */
package controllers

import "fmt"

type Article struct {
	Base
}

func (c *Article) Index() {
	fmt.Println(c.TplName)
}
