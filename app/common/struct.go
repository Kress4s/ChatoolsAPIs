package common

// 扫码成功的check接口返回值
type RestQRcode struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    struct {
		Alias      string `json:"alias"`
		HeadImgURL string `json:"head_image_url"`
		NickName   string `json:"nick_name"`
		Token      string `json:"token"`
		WXID       string `json:"wx_id"`
		Status     string `json:"status"`
	} `json:"data"`
}

// WXUser 微信联系人的结构体
type WXUser struct {
	WXID              string `json:"id"`
	NickName          string `json:"nick_name"`
	AliasName         string `json:"alias_name"`
	Sex               int    `json:"sex"`
	HeadBigImageURL   string `json:"head_big_image_url"`
	HeadSmallImageURL string `json:"head_small_image_url"`
	Inviter           string `json:"inviter"`
	Friend            bool   `json:"friend"`
}

// AnnounceMent 设置群公告内容
type AnnounceMent struct {
	Announcement string `json:"announcement"`
	GroupID      string `json:"group_id"`
}

// ProtoMessage 底层协议推送的消息
type ProtoMessage struct {
	FromUserName struct {
		Str string `json:"str"`
	} `json:"from_user_name"` //
	ToUserName struct {
		Str string `json:"str"`
	} `json:"to_user_name"` //
	MsgType int `json:"msg_type"` // 消息类型 10002(踢人、加人的消息类型(xml))
	Content struct {
		Str string `json:"str"`
	} `json:"content"` // 内容(我发：{"str":"程序监控你"}；别人发：{"str":"aaaa520jj:\nG吐总冠军"})
	Status      int    `json:"status"`       //貌似群的消息都是
	CreateTime  int    `json:"create_time"`  //消息时间戳
	MsgSource   string `json:"msg_source"`   // ?
	PushContent string `json:"push_content"` //提示消息(聊天输入框提示) (别人发有这个字段，我发没有这个字段)
}
