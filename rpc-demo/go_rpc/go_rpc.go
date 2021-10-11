package go_rpc

type Rpc struct { // 客户端要传输的数据，也是服务端要接收的数据
	Name string
	Age  int
}

type RpcReply struct { // 服务端要返回的数据，也是客户端想要获得的结果
	Systeminfo string
}
