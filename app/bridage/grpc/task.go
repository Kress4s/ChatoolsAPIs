package grpc

import (
	"ChatoolsAPIs/app/bridage/constant"
	pb "ChatoolsAPIs/app/bridage/grpc/proto"
	bridageModels "ChatoolsAPIs/app/bridage/models"
	"ChatoolsAPIs/app/common"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
)

// BotWorker ... 后期优化
type BotWorker struct {
	// Conn  *grpc.ClientConn
	// Lock  sync.Mutex
	Token        string
	CallBackAddr string
}

const (
	// Address grpc连接地址
	Address string = constant.GRPC_BASE_URL
)

// var conn *grpc.ClientConn // 一个连接
// var lock sync.Mutex

// GetConnInstance 获取连接
// func GetConnInstance() (*grpc.ClientConn, error) {
// 	var err error
// 	lock.Lock()
// 	defer lock.Unlock()
// 	if conn == nil {
// 		if conn, err = grpc.Dial(Address, grpc.WithInsecure()); err != nil {
// 			logs.Error("create grpc conn failed, err is ", err.Error())
// 			return nil, err
// 		}
// 	}
// 	return conn, nil
// }

// NewBotWorker ...
func NewBotWorker() *BotWorker {
	return new(BotWorker)
}

// PrepareParams 预置参数
func (c *BotWorker) PrepareParams(token, callBackAddr string) {
	c.CallBackAddr = callBackAddr
	c.Token = token
}

// Run 开始监听
func (c *BotWorker) Run() {
	var err error
	// ctx, cancle := context.WithCancel(context.Background())
	defer func() {
		/*
			TODO:
			异常退出
			1. 退出当前goroutine
			2. 更改数据库机器人的状态
			3. 记录日志(微信号、掉线时间)
			4. 通过websoket方式通知web端掉线的微信号
		*/
		// cancle() //通知所有的goroutine退出
		if err = bridageModels.UpdateBotLoginStatusByToken(c.Token); err == nil {
			logs.Info("%s has offlined, please check it to relogin", c.Token)
		}
		// wetsocket 通知前端

	}()
	var conn *grpc.ClientConn
	if conn, err = grpc.Dial(Address, grpc.WithInsecure()); err != nil {
		logs.Error("create grpc conn failed, err is ", err.Error())
		return
	}
	grpcClient := pb.NewRockRpcServerClient(conn)
	req := pb.StreamRequest{
		Token: &c.Token,
	}
	res, verr := grpcClient.Sync(context.Background(), &req)
	if verr != nil {
		log.Fatalf("Call Route err: %v", verr)
	}
	for {
		var message common.ProtoMessage
		fmt.Println("开始监控")
		response, verr := res.Recv()
		if verr != nil {
			break
		}
		if err = json.Unmarshal([]byte(*response.Payload), &message); err == nil {
			// 协程推送
			go PushMessage(message, c.CallBackAddr)
		} else {
			logs.Error("json Unmarshal meaasge failed, err is ", err.Error())
			break
		}
	}
}

// PushMessage ...
func PushMessage(message common.ProtoMessage, callBackAddr string) {
	if !strings.Contains(callBackAddr, "http://") {
		callBackAddr = "http://" + callBackAddr
	}
	httplib.Post(callBackAddr).JSONBody(&message)
	verr := recover()
	if verr != nil {
		logs.Error(verr)
	}
}
