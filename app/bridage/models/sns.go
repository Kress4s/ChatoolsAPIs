package models

import (
	"ChatoolsAPIs/app/bridage/constant"
	"fmt"

	"github.com/York-xia/tools/curd/common"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

// Sns ...
type Sns struct {
	ID         int64  `json:"id"`
	WXID       string `json:"wx_id"`
	NickName   string `json:"nick_name"`
	XML        string `json:"xml"`
	CreateTime int    `json:"create_time"`
	LikeCount  int    `json:"like_count"` //点赞数
	SnsLikes   []struct {
		WXID       string `json:"wx_id"`
		CreateTime int    `json:"create_time"`
	} `json:"sns_likes"` // 点赞人的信息
	CommentCount int // 评论数
	SnsComments  []struct {
		CommentID      int    `json:"comment_id"`
		ReplyCommentID int    `json:"reply_comment_id"`
		WXID           string `json:"wx_id"`
		Content        string `json:"content"`
		CreateTime     int    `json:"create_time"`
	}
}

// SnsList ...
type SnsList struct {
	FirstPageMd5 string `json:"first_page_md5"`
	SnsObjects   []Sns  `json:"sns_objects"`
}

// ImageSns 发送图片的朋友圈请求体
type ImageSns struct {
	AtList    []string `json:"at_list"`
	BlackList []string `json:"black_list"`
	Content   string   `json:"content"`
	ImageURL  []string `json:"image_url"`
}

// TextSns 发送文字的朋友圈请求体
type TextSns struct {
	AtList    []string `json:"at_list"`
	BlackList []string `json:"black_list"`
	Content   string   `json:"content"`
}

// VideoSns 发送图片的朋友圈请求体
type VideoSns struct {
	AtList    []string `json:"at_list"`
	BlackList []string `json:"black_list"`
	Content   string   `json:"content"`
	VideoURL  []string `json:"video_url"`
}

// ListMySns ...
func ListMySns(firstPageMd5, maxID, token string) (l interface{}, err error) {
	var sns SnsList
	var v = common.StandardRestResult{}
	v.Data = sns
	// 这里前两
	if err = httplib.Get(constant.SNS_LIST_ME).Header(constant.H_AUTHORIZATION, token).Param("first_page_md5", firstPageMd5).Param("max_id", maxID).ToJSON(&v); err != nil {
		logs.Error("ListMySns: ToJSON sns failed, err is ", err.Error())
		return
	}
	return v, nil
}

// ListMyFriend ...
func ListMyFriend(firstPageMd5, maxID, WXID, token string) (l interface{}, err error) {
	var sns SnsList
	var v = common.StandardRestResult{}
	v.Data = sns
	// 这里前两个参数，第一次传空未测试
	if err = httplib.Get(constant.SNS_LIST_ME).Header(constant.H_AUTHORIZATION, token).Param("wx_id", WXID).Param("first_page_md5", firstPageMd5).Param("max_id", maxID).ToJSON(&v); err != nil {
		logs.Error("ListMyFriendSns: ToJSON sns failed, err is ", err.Error())
		return
	}
	return v, nil
}

// SendImageSns ...
func SendImageSns(m ImageSns, token string) (ret interface{}, err error) {
	var sns Sns
	var v = common.StandardRestResult{}
	v.Data = sns
	req, verr := httplib.Post(constant.SNS_SEND_IMAGE).Header(constant.H_AUTHORIZATION, token).JSONBody(&m)
	if verr != nil {
		logs.Error("[%s] SendImageSns Content[%s] failed, err is ", token, m.Content, err.Error())
		return nil, verr
	}
	if err = req.ToJSON(&v); err != nil {
		logs.Error("[%s] SendImageSns ToJSON failed, err is ", token, err.Error())
		return
	}
	if v.Code == 0 {
		ret = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return
}

// SendTextSns ...
func SendTextSns(m TextSns, token string) (ret interface{}, err error) {
	var sns Sns
	var v = common.StandardRestResult{}
	v.Data = sns
	req, verr := httplib.Post(constant.SNS_SEND_TEXT).Header(constant.H_AUTHORIZATION, token).JSONBody(&m)
	if verr != nil {
		logs.Error("[%s] SendImageSns Content[%s] failed, err is ", token, m.Content, err.Error())
		return nil, verr
	}
	if err = req.ToJSON(&v); err != nil {
		logs.Error("[%s] SendImageSns ToJSON failed, err is ", token, err.Error())
		return
	}
	if v.Code == 0 {
		ret = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return
}

// SendVideoSns ...
func SendVideoSns(m VideoSns, token string) (ret interface{}, err error) {
	var sns Sns
	var v = common.StandardRestResult{}
	v.Data = sns
	req, verr := httplib.Post(constant.SNS_SEND_VIDEO).Header(constant.H_AUTHORIZATION, token).JSONBody(&m)
	if verr != nil {
		logs.Error("[%s] SendImageSns Content[%s] failed, err is ", token, m.Content, err.Error())
		return nil, verr
	}
	if err = req.ToJSON(&v); err != nil {
		logs.Error("[%s] SendImageSns ToJSON failed, err is ", token, err.Error())
		return
	}
	if v.Code == 0 {
		ret = v.Data
	} else {
		err = fmt.Errorf(v.Message)
	}
	return
}
