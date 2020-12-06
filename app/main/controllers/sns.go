package controllers

import (
	"ChatoolsAPIs/app/bridage/constant"
	bridageModels "ChatoolsAPIs/app/bridage/models"
	"encoding/json"
	"fmt"

	"github.com/York-xia/tools/curd/common"
)

// SnsControllers ...
type SnsControllers struct {
	common.BaseController
}

// ListMe 我的朋友圈列表
// @router /list/me [get]
func (c *SnsControllers) ListMe() {
	var l interface{}
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandardRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandardRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	var firstPageMd5, maxID string
	firstPageMd5, maxID = c.GetString("first_page_md5"), c.GetString("maxmax_id")
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.ListMySns(firstPageMd5, maxID, token)
}

// ListMyFriendSns 我的朋友的朋友圈列表
// @router /list/friend [get]
func (c *SnsControllers) ListMyFriendSns() {
	var l interface{}
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandardRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandardRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	var firstPageMd5, maxID, WXID string
	if WXID = c.GetString("wx_id"); WXID == "" {
		err = fmt.Errorf("wx_id cant be null")
		return
	}
	firstPageMd5, maxID = c.GetString("first_page_md5"), c.GetString("maxmax_id")
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.ListMyFriend(firstPageMd5, maxID, WXID, token)
}

// SendImageAndTextSns 发送图文朋友圈
// @router /send/image [post]
func (c *SnsControllers) SendImageAndTextSns() {
	var l interface{}
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandardRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandardRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	var v bridageModels.ImageSns
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.SendImageSns(v, token)
}

// SendTheTextSns 发送文字朋友圈
// @router /send/text [post]
func (c *SnsControllers) SendTheTextSns() {
	var l interface{}
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandardRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandardRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	var v bridageModels.TextSns
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.SendTextSns(v, token)
}

// SendTheVideoSns 发送视频&文字朋友圈
// @router /send/video [post]
func (c *SnsControllers) SendTheVideoSns() {
	var l interface{}
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandardRestResult{Code: 0, Message: "ok", Data: l}
		} else {
			c.Data["json"] = common.StandardRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	var v bridageModels.VideoSns
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.SendVideoSns(v, token)
}
