package controllers

import (
	"ChatoolsAPIs/app/bridage/constant"
	bridageModels "ChatoolsAPIs/app/bridage/models"
	"ChatoolsAPIs/app/common"
	"ChatoolsAPIs/app/main/models"
	"encoding/json"
	"fmt"
)

// MessageController ...
type MessageController struct {
	common.BaseController
}

// SendTheText ...
// @router /send/text [post]
func (c *MessageController) SendTheText() {
	var l interface{}
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	var message bridageModels.SendTextMessage
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &message); err != nil {
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.SentTextMessage(message, token)
}

// SendTheImage ...
// @router /send/image [get]
func (c *MessageController) SendTheImage() {
	var l interface{}
	var to, URL string
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	if to, URL = c.GetString("to"), c.GetString("url"); to == "" || URL == "" {
		err = fmt.Errorf("to or url cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.SentImageMessage(to, URL, token)
}

// SendTheVideo ...
// @router /send/video [get]
func (c *MessageController) SendTheVideo() {
	var l interface{}
	var to, URL string
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	if to, URL = c.GetString("to"), c.GetString("url"); to == "" || URL == "" {
		err = fmt.Errorf("to or url cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.SentVideoMessage(to, URL, token)
}

// SyncRecieveMessage 接受微信号消息
// @router /sync/message/pushstream
func (c *MessageController) SyncRecieveMessage() {
	var l interface{}
	var callbackAddr string
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	if callbackAddr = c.GetString("callbackAddr"); callbackAddr == "" {
		err = fmt.Errorf("callbackAddr cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	err = models.SyncRecieveMessageStream(callbackAddr, token)
}
