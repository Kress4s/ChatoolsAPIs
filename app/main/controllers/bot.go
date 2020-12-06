package controllers

import (
	"ChatoolsAPIs/app/bridage/constant"
	bridageModels "ChatoolsAPIs/app/bridage/models"

	"github.com/York-xia/tools/curd/common"
)

// BotController ...
type BotController struct {
	common.BaseController
}

// URLMapping ...
func (c *BotController) URLMapping() {
	c.Mapping("GetQR", c.GetQR)
}

// GetQR ...
// @router /getqrcode
func (c *BotController) GetQR() {
	var v interface{}
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.RestResult{Code: 0, Message: "ok", Data: v}
		} else {
			c.Data["json"] = common.RestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	token := c.Ctx.Input.Header("Authorization")
	if token == "" {
		err = constant.ErrNilAuthorization
		return
	}
	v, err = bridageModels.GetQRCode(token)
}

// GetMyInfo ...
// @router /profile [get]
func (c *BotController) GetMyInfo() {
	var l interface{}
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.RestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.RestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	token := c.Ctx.Input.Header("Authorization")
	l, err = bridageModels.GetMyProfile(token)
}
