package models

import (
	"ChatoolsAPIs/app/bridage/constant"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/York-xia/tools/curd/common"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

// SendTextMessage 发送文字信息
type SendTextMessage struct {
	At      []string `json:"at"`      // 发送者
	To      string   `json:"to"`      // 接受者
	Content string   `json:"content"` // word content
}

// SentTextMessage ...
func SentTextMessage(m SendTextMessage, token string) (l interface{}, err error) {
	res, verr := httplib.Post(constant.SEND_TEXT).Header(constant.H_AUTHORIZATION, token).JSONBody(&m)
	if verr != nil {
		logs.Error("[%+v] send message to [%s] faield, err is %s", m.At[0], m.To, err.Error())
		return nil, verr
	}
	var response common.StandardRestResult
	if err = res.ToJSON(&response); err != nil {
		logs.Error("ToJSON: send text interface response failed, err is", err.Error())
		return nil, err
	}
	if response.Code == 0 {
		l = response.Data
	} else {
		err = fmt.Errorf(response.Message)
	}
	return
}

// SentImageMessage ...
func SentImageMessage(to, URL, token string) (l interface{}, err error) {
	resp, verr := httplib.Get(constant.SEND_IMAGE).Header(constant.H_AUTHORIZATION, token).Param("to", to).Param("url", URL).DoRequest()
	if verr != nil {
		logs.Error("send image[%s] to[%s] failed, err is ", URL, to)
		return nil, verr
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	var response common.StandardRestResult
	if err = json.Unmarshal(body, &response); err != nil {
		logs.Error("SendImage: json Unmarshal failed, err is ", err.Error())
		return nil, err
	}
	// 目前地底层发送成功和失败code都是0，没明确提示
	if response.Code != 0 {
		err = fmt.Errorf("%s", response.Message)
		logs.Error("send message[%s] to receiver[%s] failed, err is ", URL, to, err.Error())
		return nil, err
	}
	if response.Code == 0 {
		l = response.Data
	} else {
		err = fmt.Errorf(response.Message)
	}
	return
}

// SentVideoMessage ...
func SentVideoMessage(to, URL, token string) (l interface{}, err error) {
	resp, verr := httplib.Get(constant.SEND_VIDEO).Header(constant.H_AUTHORIZATION, token).Param("to", to).Param("url", URL).DoRequest()
	if verr != nil {
		logs.Error("send image[%s] to[%s] failed, err is ", URL, to)
		return nil, verr
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	var response common.StandardRestResult
	if err = json.Unmarshal(body, &response); err != nil {
		logs.Error("SendImage: json Unmarshal failed, err is ", err.Error())
		return nil, err
	}
	// 目前地底层发送成功和失败code都是0，没明确提示
	if response.Code != 0 {
		err = fmt.Errorf("%s", response.Message)
		logs.Error("send message[%s] to receiver[%s] failed, err is ", URL, to, err.Error())
		return nil, err
	}
	if response.Code == 0 {
		l = response.Data
	} else {
		err = fmt.Errorf(response.Message)
	}
	return
}
