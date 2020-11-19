package models

import (
	"ChatoolsAPIs/app/bridage/grpc"
)

// SyncRecieveMessageStream 消息推送
func SyncRecieveMessageStream(callbackAddr, token string) (err error) {
	// 开启监听此微信号
	botWork := grpc.NewBotWorker()
	botWork.PrepareParams(token, callbackAddr)
	// goroutine 监听
	go botWork.Run()
	return
}
