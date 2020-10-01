package main

import (
	"IPFS/client/clientProcess"
	"fmt"
	"net"
	"os"
)

func main()  {
	var key int
	var loop = true
	var conn net.Conn
	var code int
	for  loop {
		fmt.Println("---------------------------------------------------------------------")
		fmt.Println("\t\t 请输入登录身份")
		fmt.Println("\t\t1. :A")
		fmt.Println("\t\t2. :B")
		fmt.Scanf("%d\n", &key)
		if key ==1 || key == 2{
			conn, code = clientProcess.Login(key)
			break
		}else{
			fmt.Println("无效输入")
		}

		}
		if code == 200{
			for{
				fmt.Println("--------------已经连接上服务器，请选择服务------------------ ")
				fmt.Println("\t\t1. :上传消息")
				fmt.Println("\t\t2. :下载消息")
				fmt.Println("\t\t3. :退出系统")
				fmt.Scanf("%d\n", &key)
				switch key {
				case 1:
					clientProcess.UpLoad(conn)
				case 2:
					clientProcess.DownLoad(conn)
				default:
					os.Exit(0)
				}
			}
		}

	}