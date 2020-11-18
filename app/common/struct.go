package common

//RestResult Rest接口返回值
type RestResult struct {
	Code    int         // 0 表示成功，其他失败
	Message string      // 错误信息
	Data    interface{} // 数据体
}

// 标准的 rest 返回接口，字符小写化
type StandRestResult struct {
	Code    int         `json:"code"` // 0 表示成功，其他失败
	Message string      `json:"msg"`  // 错误信息
	Data    interface{} `json:"data"` // 数据体
}

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
