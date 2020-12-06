package models

import (
	"ChatoolsAPIs/app/bridage/constant"
	"fmt"

	"github.com/York-xia/tools/curd/common"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

// Contact 联系人
type Contact struct {
	WXID           string `json:"id"`
	BigHeadImage   string `json:"head_big_image_url"`
	SmallHeadImage string `json:"head_small_image_url"`
	NickName       string `json:"nick_name"`
	Country        string `json:"country"`
	Province       string `json:"province"`
	City           string `json:"city"`
	Sex            bool   `json:"sex"`
	Signature      string `json:"signature"`
	Alias          string `json:"alias_name"`
	// Labels         []*Label `orm:"reverse(many)"`
}

// result contact query result
type result struct {
	CurrentWxContactSeq       int      `json:"current_wx_contact_seq"`
	CurrentChatRoomContactSeq int      `json:"current_chat_room_contact_seq"`
	IDs                       []string `json:"ids"`
}

// ContactBetch ...
func ContactBetch(query, token string) (m interface{}, err error) {
	var contact Contact
	var v = common.StandardRestResult{}
	v.Data = contact
	if err = httplib.Get(constant.CONTACT_BATCH_URL+query).Header(constant.H_AUTHORIZATION, token).ToJSON(&v); err != nil {
		logs.Error("ContactBetch: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	if v.Code == 0 {
		m = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return m, err
}

// ContactListAll ...
func ContactListAll(room_contact_seq, wx_contact_seq, token string) (l interface{}, err error) {
	var _v result
	var v = common.StandardRestResult{}
	v.Data = _v
	query := "?room_contact_seq=" + room_contact_seq + "&wx_contact_seq=" + wx_contact_seq
	if err = httplib.Get(constant.CONTACT_LIST_URL+query).Header(constant.H_AUTHORIZATION, token).ToJSON(&v); err != nil {
		logs.Error("ContactListAll: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	if v.Code == 0 {
		l = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return
}

// ContactListGroup ...
func ContactListGroup(room_contact_seq, wx_contact_seq, token string) (l interface{}, err error) {
	var _v result
	var v = common.StandardRestResult{}
	v.Data = _v
	query := "?room_contact_seq=" + room_contact_seq + "&wx_contact_seq=" + wx_contact_seq
	if err = httplib.Get(constant.CONTACT_GROUP_LIST_URL+query).Header(constant.H_AUTHORIZATION, token).ToJSON(&v); err != nil {
		logs.Error("ContactListAll: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	if v.Code == 0 {
		l = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return
}

// ContactSearch ...
func ContactSearch(keyword, token string) (l interface{}, err error) {
	var contact Contact
	var v = common.StandardRestResult{}
	v.Data = contact
	if err = httplib.Get(constant.CONTACT_SEARCH_URL).Header(constant.H_AUTHORIZATION, token).Param("keyword", keyword).ToJSON(&v); err != nil {
		logs.Error("ContactBetch: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	if v.Code == 0 {
		l = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return l, err
}
