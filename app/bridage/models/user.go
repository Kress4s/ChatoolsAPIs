package models

import (
	"ChatoolsAPIs/app/common"
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// User ...
type User struct {
	ID       int64  `orm:"auto;column(id)"`
	Account  string `orm:"size(20);column(account)"`
	PassWord string `orm:"size(20);column(password)"`
	BotNum   int    `orm:"column(bot_num)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser ...
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	if id, err = o.Insert(m); err != nil {
		logs.Error("AddLabel failed, err is ", err.Error())
		return
	}
	return
}

// GenerateToken ...
func GenerateToken(m *User) (token string, err error) {
	o := orm.NewOrm()
	defer func() {
		if err == nil {
			o.Commit()
		} else {
			o.Rollback()
		}
	}()
	o.Begin()
	// judging can get token or not
	var user = User{Account: m.Account}
	if !o.QueryTable(new(User)).Filter("Account", m.Account).Exist() {
		err = fmt.Errorf("该用户不存在")
		return "", err
	}
	if err = o.Read(&user, "Account"); err != nil {
		return "", err
	}
	var botNum int64
	if botNum, err = o.QueryTable(new(Bots)).Filter("User", m.Account).Count(); int(botNum) >= user.BotNum {
		err = fmt.Errorf("本账号token数量权限已满")
		return "", err
	}
	token = common.GenetateAuth()
	var bot = new(Bots)
	bot.Token = token
	bot.User = user.Account
	if _, err = o.Insert(bot); err != nil {
		logs.Error("insert bot failed, err is", err.Error())
		return "", err
	}
	user.BotNum++
	var num int64
	if num, err = o.Update(&user, "BotNum"); err != nil {
		logs.Error("update user failed, err is", err.Error())
		return "", err
	}
	logs.Debug("Number of User update in database:", num)
	return token, nil
}
