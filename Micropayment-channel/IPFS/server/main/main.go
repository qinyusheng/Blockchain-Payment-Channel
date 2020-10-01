package main

import (
	"IPFS/server/serverProcess"
	"fmt"
	"net"
	"os"
)

var ipfsLog *os.File

func logInit() {
	log, err := os.Create("./ipfs_log")
	ipfsLog = log
	if err != nil {
		panic(err)
	}
}

func ilog(rec string) {
	_, err := ipfsLog.WriteString(rec);
	if (err != nil) {
		panic(err)
	}
}


func processor(conn net.Conn)  {
	defer conn.Close()
	//创建一个processor实例，
	processor :=  serverProcess.Processor{
		Conn: conn,
	}
	//调用总控
	err := processor.ServerProcessor()
	if err != nil{
		fmt.Println("客户端协程错误 err=", err)
	}
	return
}

func main()  {
	fmt.Println("服务器在8889端口监听.....")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil{
		fmt.Println("net.listen err=", err)
		return
	}
	for{
		fmt.Println("等待客户端连接........")
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("listen.Accept err=", err)
		}
		go processor(conn)
	}

}