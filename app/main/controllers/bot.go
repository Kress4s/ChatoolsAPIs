package controllers

import (
	"ChatoolsAPIs/app/bridage/common"
)

type BotController struct {
	common.BaseController
}

func (c *BotController) GetQR() {
	// var err error
	token := c.Ctx.Input.Header("Authorization")
	if token == "" {
		// err = fmt.Errorf("tokon is nil")
		return
	}
	return
}
