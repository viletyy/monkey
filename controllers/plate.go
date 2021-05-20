/*
 * @Date: 2021-04-25 21:09:27
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:37:04
 * @FilePath: /monkey/controllers/plate.go
 */
package controllers

import (
	"github.com/viletyy/monkey/global"
	"github.com/viletyy/monkey/model"
)

type Plate struct {
	Base
}

func (c *Plate) Index() {
	plateSearch := global.Search{
		Maps: make(map[string]interface{}),
		PageInfo: global.PageInfo{
			Page:     1,
			PageSize: 10,
		},
	}
	plateResult, plateErr := model.GetPlates(&plateSearch)
	if plateErr != nil {
		c.FlashError("加载板块失败")
	}
	c.Data["Plates"] = plateResult.List
}

func (c *Plate) Show() {
	var plateId int
	err := c.Ctx.Input.Bind(&plateId, ":id")
	if err != nil {
		c.ErrorHandler(err)
	}
	if dbPlate, err := model.GetPlateById(plateId); err != nil {
		c.ErrorHandler(err)
	} else {
		c.Data["Plate"] = dbPlate
		maps := make(map[string]interface{})
		maps["plate_id"] = dbPlate.ID
		// TODO 关键词搜索
		search := global.Search{
			Maps:     maps,
			PageInfo: c.PageInfo(),
		}

		detailResult, detailErr := model.GetDetails(&search)
		if detailErr != nil {
			c.FlashError("加载列表失败")
		}

		c.ResponseWithResult(&detailResult)
	}
}
