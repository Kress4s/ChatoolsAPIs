package common

//RestResult Rest接口返回值
type RestResult struct {
	Code    int         // 0 表示成功，其他失败
	Message string      // 错误信息
	Data    interface{} // 数据体
}

// 标准的 rest 返回接口，字符小写化
type StandRestResult struct {
	Code    int         `json:"code"`    // 0 表示成功，其他失败
	Message string      `json:"message"` // 错误信息
	Data    interface{} `json:"data"`    // 数据体
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