package main

import (
	"fmt"
	"net"
	"os"
	"net1"
        "IPFS/client/clientProcess"
)

var iPaddress string = "10.0.0.1:8888"
var conn net.Conn
var ipfs net.Conn
var testTimes int = 1000

func BuyTest(conn net.Conn) {
    var num int
    fmt.Print("Please enter the times of test you want: ")
    fmt.Scan(&testTimes)
    fmt.Print("Please enter the range of data id: ")
    fmt.Scan(&num)

	for i:=0; i<testTimes; i ++ {
		addr, SPub := net1.BuyData(conn, i%num, 1)
		if (addr == "") {
			fmt.Println("Buy Fail")
			break;
		} else {
			str := net1.DownLoad(ipfs, addr, SPub)
			fmt.Println("date content:", str);
		}
	}
}

func main() {
	var code int
	var err error
	iPaddress = os.Args[2]+":8888"
	conn ,err = net.Dial("tcp", iPaddress)

	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()

	ipfs, code = clientProcess.LoginWithIP(1, os.Args[4])

        if (code != 200) {
                panic("1号连接失败")
        }
      
        net1.ClientInit()
	net1.CreateBond(conn)
	
	BuyTest(conn)

/*
	for true {
		fmt.Println("Please input date id (>0) to continue buy date");
		var id int
		fmt.Scanf("%d", &id);
		
		if (id < 0) {
			fmt.Println("Transaction Complete!");
			break;
		}
		
		addr, SPub := net1.BuyData(conn, id, 1)
		if (addr == "") {
			fmt.Println("Buy Fail")
			break;
		} else {
			str := net1.DownLoad(ipfs, addr, SPub)
        		fmt.Println("date content:", str);
		}
	}
*/
	
	net1.ClientClose()
	fmt.Println("Client Close")
}


