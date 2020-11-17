package controllers

import (
	"ChatoolsAPIs/app/bridage/common"
	bridageModels "ChatoolsAPIs/app/bridage/models"
	"encoding/json"
)

type UserController struct {
	common.BaseController
}

// Register ...
func (c *UserController) Register() {
	var err error
	var v bridageModels.User
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.RestResult{Code: 0, Message: "ok", Data: v}
		} else {
			c.Data["json"] = common.RestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		return
	}
	_, err = bridageModels.AddUser(&v)
}

// Register ...
// @router /token [post]
func (c *UserController) GenerateToken() {
	var token string
	var err error
	defer func() {
		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = common.RestResult{Code: 0, Message: "ok", Data: token}
		} else {
			c.Data["json"] = common.RestResult{Code: -1, Message: err.Error()}
		}
		c.ServeJSON()
	}()
	var user bridageModels.User
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		return
	}
	token, err = bridageModels.Generate(&user)
}
