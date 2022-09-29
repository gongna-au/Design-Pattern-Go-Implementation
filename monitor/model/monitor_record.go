package model

import "sync/atomic"

type Type string

const (
	RecvReq  Type = "recv_req"  // 接收请求
	RecvResp Type = "recv_resp" // 接收响应
	SendReq  Type = "send_req"  // 发送请求
	SendResp Type = "send_resp" // 发送响应
)

// id生成器
var recordId int32 = 0

// MonitorRecord 监控记录
type MonitorRecord struct {
	//唯一标志独一无二的监控记录
	//自增
	Id        int
	Endpoint  string
	Type      Type
	Timestamp int64
}

func NewMonitoryRecord() *MonitorRecord {
	return &MonitorRecord{
		Id: int(atomic.AddInt32(&recordId, 1)),
	}
}
