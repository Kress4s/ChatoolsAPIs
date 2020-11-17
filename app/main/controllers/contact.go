package controllers

import (
	"ChatoolsAPIs/app/bridage/constant"
	bridageModels "ChatoolsAPIs/app/bridage/models"
	"ChatoolsAPIs/app/common"
	"fmt"
	"strings"
)

// ContactController ...
type ContactController struct {
	common.BaseController
}

// URLMapping ...
func (c *ContactController) URLMapping() {
	c.Mapping("Betch", c.Betch)
}

// Betch 批量的到联系人的详细信息
// @router /betch [get]
func (c *ContactController) Betch() {
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
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	ids := c.GetStrings("ids")
	if len(ids) == 0 {
		err = fmt.Errorf("ids must have one")
		return
	}
	query := "?ids=" + strings.Join(ids, "&ids=")
	l, err = bridageModels.ContactBetch(query, token)
}

// ListAll 得到所有联系人(包括群)的ID列表
// @router /list/all [get]
func (c *ContactController) ListAll() {
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
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	var room_contact_seq, wx_contact_seq string
	if room_contact_seq = c.GetString("room_contact_seq"); room_contact_seq == "" {
		room_contact_seq = "0"
	}
	if wx_contact_seq = c.GetString("wx_contact_seq"); wx_contact_seq == "" {
		wx_contact_seq = "0"
	}
	l, err = bridageModels.ContactListAll(room_contact_seq, wx_contact_seq, token)
}

// ListGroup 得到群ID列表
// @router /list/group [get]
func (c *ContactController) ListGroup() {
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
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	var room_contact_seq, wx_contact_seq string
	if room_contact_seq = c.GetString("room_contact_seq"); room_contact_seq == "" {
		room_contact_seq = "0"
	}
	if wx_contact_seq = c.GetString("wx_contact_seq"); wx_contact_seq == "" {
		wx_contact_seq = "0"
	}
	l, err = bridageModels.ContactListGroup(room_contact_seq, wx_contact_seq, token)
}

// Search 搜索，用户添加好友
// @router /search [get]
func (c *ContactController) Search() {
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
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	search := c.GetString("keyword")
	if search == "" {
		err = fmt.Errorf("keyword cant be null")
		return
	}
	l, err = bridageModels.ContactSearch(search, token)
}
