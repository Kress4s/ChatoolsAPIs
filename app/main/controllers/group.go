package controllers

import (
	"ChatoolsAPIs/app/bridage/constant"
	bridageModels "ChatoolsAPIs/app/bridage/models"
	"ChatoolsAPIs/app/common"
	"encoding/json"
	"fmt"
	"strings"
)

// GroupController ...
type GroupController struct {
	common.BaseController
}

// Create 创建群组，目前有问题，未调
// @router /create [get]
func (c *GroupController) Create() {
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
	members := c.GetStrings("members")
	if len(members) == 0 {
		err = fmt.Errorf("ids must have one")
		return
	}
	query := "?ids=" + strings.Join(members, "&ids=")
	l, err = bridageModels.GroupCreate(query, token)
}

// GetGroupInfoA 获取群公告内容
// @router /get/detail [get]
func (c *GroupController) GetGroupInfoA() {
	var groupID string
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
	if groupID = c.GetString("group_id"); groupID == "" {
		err = fmt.Errorf("group_id cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.GetGroupAnnounce(groupID, token)
}

// GetGroupInfoB 获取群信息，无公告内容
// @router /get/info [get]
func (c *GroupController) GetGroupInfoB() {
	var groupID string
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
	if groupID = c.GetString("group_id"); groupID == "" {
		err = fmt.Errorf("group_id cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.GetGroupInfo(groupID, token)
}

// GetGroupMembers 获取群成员信息
// @router /get/members [get]
func (c *GroupController) GetGroupMembers() {
	var groupID string
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
	if groupID = c.GetString("group_id"); groupID == "" {
		err = fmt.Errorf("group_id cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	l, err = bridageModels.GetGroupMembers(groupID, token)
}

// AddGroupMember 增加群成员
// @router /members/add [get]
func (c *GroupController) AddGroupMember() {
	var groupID string
	var members []string
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandRestResult{Code: 0, Message: "ok"}
		} else {
			c.Data["json"] = common.StandRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	if groupID, members = c.GetString("group_id"), c.GetStrings("members"); groupID == "" || len(members) == 0 {
		err = fmt.Errorf("group_id or members cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	err = bridageModels.AddMember(groupID, token, members)
}

// DeleteGroupMember 删除群成员[群主]
// @router /members/delete [get]
func (c *GroupController) DeleteGroupMember() {
	var groupID string
	var members []string
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandRestResult{Code: 0, Message: "ok"}
		} else {
			c.Data["json"] = common.StandRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	if groupID, members = c.GetString("group_id"), c.GetStrings("members"); groupID == "" || len(members) == 0 {
		err = fmt.Errorf("group_id or members cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	err = bridageModels.DelMember(groupID, token, members)
}

// QuitGroup 退群
// @router /quit [get]
func (c *GroupController) QuitGroup() {
	var groupID string
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandRestResult{Code: 0, Message: "ok"}
		} else {
			c.Data["json"] = common.StandRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	if groupID = c.GetString("group_id"); groupID == "" {
		err = fmt.Errorf("group_id or members cant be null")
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	err = bridageModels.QuitFromGroup(groupID, token)
}

// SetAnnounceMent 设置群公告
// @router /set/announcement [post]
func (c *GroupController) SetAnnounceMent() {
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.StandRestResult{Code: 0, Message: "ok"}
		} else {
			c.Data["json"] = common.StandRestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	var announce common.AnnounceMent
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &announce); err != nil {
		return
	}
	token := c.Ctx.Input.Header(constant.H_AUTHORIZATION)
	err = bridageModels.SetAnnounce(announce, token)
}
