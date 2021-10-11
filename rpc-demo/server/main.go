package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"rpc-demo/go_rpc"
	"runtime"
)

type TakeRpc struct{}

func (t TakeRpc) GetSystem(arg *go_rpc.Rpc, result *go_rpc.RpcReply) error {
	fmt.Printf("客户端发送了： %s %d\n", arg.Name, arg.Age)
	// 返回给服务端的
	result.Systeminfo = runtime.GOOS
	return nil
}

func main() {
	tr := new(TakeRpc)
	// 注册RPC可以一次性注册多个RPC服务，可以用FOR注册多个结构体，用web api的话就是注册多个控制器
	err := rpc.Register(tr)
	if err != nil {
		log.Fatalln("注册方法时出现的错误", err)
	}
	// 实际上RPC的底层也是基于TCP链接的，开放8081供客户端连入
	l, err := net.Listen("tcp", ":8081")
	// 记得关闭
	defer l.Close()
	if err != nil {
		fmt.Println("监听失败：", err)
	}

	for {
		var conn net.Conn
		conn, err = l.Accept()
		if err != nil {
			log.Fatalln("创建连接失败：", err)
		}
		go rpc.ServeConn(conn)
	}
}
