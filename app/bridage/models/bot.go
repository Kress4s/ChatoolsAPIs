package models

import (
	"ChatoolsAPIs/app/bridage/common"
	"ChatoolsAPIs/app/bridage/constant"
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
	"time"
)

type Bots struct {
	ID    int64  `orm:"auto;column(id)"`
	WXID  string `orm:"size(30);column(wx_id)"`
	Token string `orm:"size(50);column(token)"`
	User  string `orm:"size(30);column(user)"`
}

func init() {
	orm.RegisterModel(new(Bots))
}

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
	var restBody common.StandRestResult
	if err = json.Unmarshal(body, &restBody); err != nil {
		logs.Error("json.Unmarshal qrcode failed, err is ", err.Error())
		return
	}
	// 等待扫码接口（等待5分钟）
	ctx, _ := context.WithTimeout(context.Background(), 5*60*time.Second)
	go WaitScanCodeAndRecodeBotInfo(ctx, &bot)
	return restBody, nil
}

func WaitScanCodeAndRecodeBotInfo(ctx context.Context, bot *Bots) {
	var restBody common.RestQRcode
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
				if _, err = o.Update(bot, "WXID"); err == nil {
					logs.Debug("WaitScanCodeAndRecodeBotInfo: bot update WX_ID in database")
				}
				break L1
			}
		}
		// 等待用户操作
		time.Sleep(2 * time.Second)
	}
}
