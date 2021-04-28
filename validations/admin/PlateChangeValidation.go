/*
 * @Date: 2021-04-28 22:43:59
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-28 22:50:15
 * @FilePath: /monkey/validations/admin/PlateChangeValidation.go
 */
package admin

type PlateChangeValidation struct {
	Name             string `form:"name" valid:"Required"`
	ArticleRecommend bool   `form:"article_recommend"`
	NewsRecommend    bool   `form:"news_recommend"`
}
