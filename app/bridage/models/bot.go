package models

import (
	"ChatoolsAPIs/app/bridage/constant"
	localCommon "ChatoolsAPIs/app/common"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/York-xia/tools/curd/common"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Bots ...
type Bots struct {
	ID          int64  `orm:"auto;column(id)"`
	WXID        string `orm:"size(30);column(wx_id)"`
	LoginStatus int    `orm:"column(login_status)"`
	Token       string `orm:"size(50);column(token)"`
	User        string `orm:"size(30);column(user)"`
}

func init() {
	orm.RegisterModel(new(Bots))
}

// GetQRCode ...
func GetQRCode(token string) (ret interface{}, err error) {
	o := orm.NewOrm()
	var resp *http.Response
	var bot = Bots{Token: token}
	//if err = o.QueryTable(new(Bots)).Filter("Token", token).Filter("Token", token).One(&bot); {
	//	err = fmt.Errorf("Token is Invalid")
	//	return nil, err
	//}
	err = o.Read(&bot, "Token")
	// 未授权token
	if err == orm.ErrNoRows {
		err = fmt.Errorf("Token is Invalid")
		return nil, err
	}
	// 正确授权的token
	/*
		1. 判断token是新token还是老token
		2. 不同的拿二维码的方式
	*/
	if bot.WXID == "" {
		// 新token, 新的微信号
		if resp, err = httplib.Get(constant.LOGIN_QRCODE_URL).Header(constant.H_AUTHORIZATION, token).DoRequest(); err != nil {
			logs.Error("get URL[%s] failed, err is ", err.Error())
			return
		}
	} else {
		// 登录过
		if resp, err = httplib.Get(constant.LOGIN_QRCODE_URL).Param(constant.P_WXID, bot.WXID).Header(constant.H_AUTHORIZATION, token).DoRequest(); err != nil {
			logs.Error("get URL[%s] failed, err is ", err.Error())
			return
		}
	}
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("get URL[%s] body failed, err is ", constant.LOGIN_QRCODE_URL, err.Error())
		return
	}
	var restBody common.StandardRestResult
	if err = json.Unmarshal(body, &restBody); err != nil {
		logs.Error("json.Unmarshal qrcode failed, err is ", err.Error())
		return
	}
	// 等待扫码接口（等待5分钟）
	ctx, _ := context.WithTimeout(context.Background(), 5*60*time.Second)
	go WaitScanCodeAndRecodeBotInfo(ctx, &bot)
	return restBody.Data, nil
}

// WaitScanCodeAndRecodeBotInfo ...
func WaitScanCodeAndRecodeBotInfo(ctx context.Context, bot *Bots) {
	var restBody localCommon.RestQRcode
	var resp *http.Response
	var err error
	o := orm.NewOrm()
L1:
	for {
		select {
		case <-ctx.Done():
			logs.Debug("waiting user operation is timeout in groutines")
			break L1
		default:
			if resp, err = httplib.Get(constant.LOGIN_CHECK_URL).Header(constant.H_AUTHORIZATION, bot.Token).DoRequest(); err != nil {
				logs.Error("get response[%s] failed, err is ", constant.LOGIN_CHECK_URL, err.Error())
				return
			}
			var body []byte
			if body, err = ioutil.ReadAll(resp.Body); err != nil {
				logs.Error("get URL[%s] body failed, err is ", constant.LOGIN_CHECK_URL, err.Error())
				return
			}
			if err = json.Unmarshal(body, &restBody); err != nil {
				logs.Error("json Unmarshal failed, err is ", err.Error())
				return
			}
			// 正常
			if restBody.Code == 0 && restBody.Data.Status == "Confirmed" {
				//var num int64
				bot.WXID = restBody.Data.WXID
				bot.LoginStatus = 1
				if _, err = o.Update(bot, "WXID", "LoginStatus"); err == nil {
					logs.Debug("WaitScanCodeAndRecodeBotInfo: bot update WX_ID in database")
				}
				break L1
			}
			logs.Info("等待扫码中...")
		}
		// 等待用户操作
		time.Sleep(2 * time.Second)
	}
}

// GetMyProfile ...
func GetMyProfile(token string) (ret interface{}, err error) {
	type ProfileUser struct {
		WXID              string `json:"wx_id"`
		NickName          string `json:"nick_name"`
		Alias             string `json:"alias"`
		Sex               int    `json:"sex"`
		Country           string `json:"country"`
		Province          string `json:"province"`
		City              string `json:"city"`
		Signature         string `json:"signature"`
		HeadBigImageURL   string `json:"big_head_img_url"`
		HeadSmallImageURL string `json:"small_head_img_url"`
	}
	var bot ProfileUser
	if err = httplib.Get(constant.WXUSER_PROFILE_URL).Header(constant.H_AUTHORIZATION, token).ToJSON(&bot); err != nil {
		logs.Error("GetMyProfile: ToJSON failed, err is ", err.Error())
		return nil, err
	}
	return bot, nil
}

// UpdateBotLoginStatusByToken ...
func UpdateBotLoginStatusByToken(token string) (err error) {
	o := orm.NewOrm()
	var bot = Bots{Token: token}
	if err = o.Read(&bot, "Token"); err == nil {
		var num int64
		bot.LoginStatus = 0
		if num, err = o.Update(&bot, "LoginStatus"); err == nil {
			logs.Debug("Number of User update in database:", num)
			return
		}
	}
	return
}
