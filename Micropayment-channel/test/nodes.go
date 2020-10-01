package main

import (
    "net1"
    "fmt"
    "net"
    "os"
    "IPFS/client/clientProcess"
    "math/rand"
    "time"
)

func  GetRandomString(l int) string {
    str := "0123456789abcdefghijklmnopqrstuvwxyz"
    bytes := []byte(str)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < l; i++ {
        result = append(result, bytes[r.Intn(len(bytes))])
    }
    return string(result)
}

func DataInit() {
    var size int
    var num int
    fmt.Print("Please enter the length of data you want: ")
    fmt.Scan(&size)
    fmt.Print("Please enter the num of test data you want: ")
    fmt.Scan(&num)

    for i:=0; i<=num; i++ {
	net1.AddDate(i, GetRandomString(size))
    }
    fmt.Println("Data init success")
}

func main()  {
        // 进行初始化
	var ipfs net.Conn
        var code int
        ipfs, code = clientProcess.LoginWithIP(2, os.Args[2])
        if (code != 200) {
                panic("2号连接失败")
        }

        net1.ServerInit()
	DataInit()

        listen,err := net.Listen("tcp","0.0.0.0:8888")
        if err != nil{
                fmt.Println("监听失败，错误：",err)
                return
        }

        fmt.Print("listen 8888: \n")
        for {
                conn,err := listen.Accept() //建立连接
                if err!= nil{
                        fmt.Println("建立连接失败，错误：",err)
                        continue
                }
                go net1.Process(conn, ipfs)    //启动一个goroutine处理连接
        }
}
