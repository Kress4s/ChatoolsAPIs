package models

import (
	"ChatoolsAPIs/app/bridage/constant"
	localCommon "ChatoolsAPIs/app/common"
	"fmt"
	"strings"

	"github.com/York-xia/tools/curd/common"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

// Group ...
type Group struct {
	GID            string `json:"wx_id"`                // json:wx_id
	NickName       string `json:"nick_name"`            //
	Owner          string `json:"owner" `               //群主
	MemberNum      int    `json:"member_num"`           //
	HeadSmallImage string `json:"head_small_image_url"` //
}

// GroupCreate ...
func GroupCreate(query, token string) (m interface{}, err error) {
	var group Group
	var v = common.StandardRestResult{}
	v.Data = group
	if err = httplib.Get(constant.GROUP_CREATE_URL+query).Header(constant.H_AUTHORIZATION, token).ToJSON(&v); err != nil {
		logs.Error("GroupCreate: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	if v.Code == 0 {
		m = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return m, err
}

// GetGroupAnnounce ...
func GetGroupAnnounce(groupID, token string) (m interface{}, err error) {
	type Info struct {
		AnnounceMent            string `json:"announcement"`
		AnnouncementEditor      string `json:"announcement_editor"`
		AnnouncementPublishTime string `json:"announcement_publish_time"`
	}
	var info Info
	var v = common.StandardRestResult{}
	v.Data = info
	if err = httplib.Get(constant.GROUP_DETAIL_URL).Header(constant.H_AUTHORIZATION, token).Param("group_id", groupID).ToJSON(&v); err != nil {
		logs.Error("GroupCreate: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	if v.Code == 0 {
		m = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return m, err
}

// GetGroupInfo ...
func GetGroupInfo(groupID, token string) (m interface{}, err error) {
	var group Group
	var v = common.StandardRestResult{}
	v.Data = group
	if err = httplib.Get(constant.GROUP_INFO_URL).Header(constant.H_AUTHORIZATION, token).Param("group_id", groupID).ToJSON(&v); err != nil {
		logs.Error("GroupCreate: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	if v.Code == 0 {
		m = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return m, err
}

// GetGroupMembers ...
func GetGroupMembers(groupID, token string) (m interface{}, err error) {
	var Members []*localCommon.WXUser
	var v = common.StandardRestResult{}
	v.Data = Members
	if err = httplib.Get(constant.GROUP_MEMBERS_URL).Header(constant.H_AUTHORIZATION, token).Param("group_id", groupID).ToJSON(&v); err != nil {
		logs.Error("GroupCreate: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	if v.Code == 0 {
		m = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return m, err
}

// AddMember ...
func AddMember(groupID, token string, members []string) (err error) {
	var v = common.StandardRestResult{}
	var query = "?group_id=" + groupID + "&members=" + strings.Join(members, "&members=")
	if err = httplib.Get(constant.GROUP_ADD_MEMBERS_URL+query).Header(constant.H_AUTHORIZATION, token).ToJSON(&v); err != nil {
		logs.Error("GroupCreate: ToJSON failed, err is ", err.Error())
		return err
	}
	if v.Code != 0 {
		err = fmt.Errorf(v.Message)
		return
	}
	return err
}

// DelMember ...
func DelMember(groupID, token string, members []string) (err error) {
	var v = common.StandardRestResult{}
	var query = "?group_id=" + groupID + "&members=" + strings.Join(members, "&members=")
	if err = httplib.Get(constant.GROUP_DEL_MEMBERS_URL+query).Header(constant.H_AUTHORIZATION, token).ToJSON(&v); err != nil {
		logs.Error("GroupCreate: ToJSON failed, err is ", err.Error())
		return err
	}
	if v.Code != 0 {
		err = fmt.Errorf(v.Message)
		return
	}
	return err
}

// QuitFromGroup ...
func QuitFromGroup(groupID, token string) (err error) {
	var v = common.StandardRestResult{}
	if err = httplib.Get(constant.GROUP_QUIT_URL).Header(constant.H_AUTHORIZATION, token).Param("group_id", groupID).ToJSON(&v); err != nil {
		logs.Error("GroupCreate: ToJSON failed, err is ", err.Error())
		return err
	}
	if v.Code != 0 {
		err = fmt.Errorf(v.Message)
		return
	}
	return err
}

// SetAnnounce ...
func SetAnnounce(announce localCommon.AnnounceMent, token string) (err error) {
	var v = common.StandardRestResult{}
	resp, verr := httplib.Post(constant.GROUP_SET_ANNOUNCE_URL).Header(constant.H_AUTHORIZATION, token).JSONBody(&announce)
	if verr != nil {
		logs.Error("SetAnnounce: ToJSON failed, err is ", err.Error())
		return err
	}
	if err = resp.ToJSON(&v); err != nil {
		logs.Error("SetAnnounce: ToJSON failed, err is ", err.Error())
	}
	if v.Code != 0 {
		err = fmt.Errorf(v.Message)
		return
	}
	return err
}
