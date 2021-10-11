package main

import (
	"log"
	"net/rpc"
	"rpc-demo/go_rpc"
)

func main() {
	// 连接服务端
	c, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln("连接服务端失败", err)
	}

	defer c.Close()
	// 初始化客户端要发送的内容
	var rpcArg go_rpc.Rpc
	// 初始化服务端要返回的内容
	var rpcReply go_rpc.RpcReply

	// 客户端发送的内容
	rpcArg.Name = "咸鱼"
	rpcArg.Age = 18

	// 直接开始发送给服务端，并获得服务端的响应
	// 第一个参数是服务端注册的RPC服务service
	if err = c.Call("TakeRpc.GetSystem", rpcArg, &rpcReply); err != nil {
		log.Fatalln("获取服务端数据错误", err)
	}
	log.Println("返回成功", rpcReply)

}
