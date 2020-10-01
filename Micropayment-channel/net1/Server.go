package net1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bitgoin/address"
	"net"
	"github.com/bitgoin/tx1"
	"strconv"
	"IPFS/client/clientProcess"
	"digital-envelope/common/box"
        "digital-envelope/common/envelope"
        "digital-envelope/common/secretbox"
	"time"
	"os"
)

// 需要放进主函数的东西
var serverKey *address.PrivateKey
var wif2 string = "92DUfNPumHzpCkKjmeqiSEDB1PU67eWbyUgYHhK9ziM7NEbqjnK"

// time value(int64, time.Now().UnixNano())
var gKeyTime int64 = 0
var ipfsUpTime int64 = 0
var ipfsUpNum int64 = 0
var txProcTime int64 = 0
var txConfTime int64 = 600 // may be a const
// size
var txSize int64 = 0
var upDataSize int64 = 0

var serverData *os.File

func sRecord(rec string) {
        _, err := serverData.WriteString(rec);
        if (err != nil) {
                panic(err)
        }
}


func ServerPrint() {
	fmt.Println("The time to generate the key:", gKeyTime, "(nm)")
	fmt.Println("The number of encrypt and upload data:", ipfsUpNum, "(nm)")

	fmt.Println("Upload data size:", upDataSize)
	
	var rec string
	rec = fmt.Sprintf("%v %v\n", ipfsUpTime, upDataSize)
	sRecord(rec)
}

type Date struct {
	id  int
	content  string
}

var dateSet []Date = []Date{}

func AddDate(id int, addr string) int {
	var date Date
	date.id = id
	date.content = addr
	dateSet = append(dateSet, date)
	return 0
}

func GetDateContent(id int) string {
	for _,d := range dateSet {
		if d.id == id {
			return d.content
		}
	}
	return ""
}

func EncryMes(date string, balanceKey string,  SPrKey string, RPuKey string) string {
	chipher, e_envelope := envelope.Seal(date, balanceKey, RPuKey, SPrKey)
	var mes Mes2
	mes.T1 = []byte(chipher)
	mes.T2 = []byte(e_envelope)
	bmes,_ := json.Marshal(mes)
	return string(bmes)
}

func UploadMes(date string, ipfsConn net.Conn) string {
	addr,_ := clientProcess.UpLoadDate(ipfsConn, date)
	return addr
}

func UploadMesById(dateID int, balanceKey string,  SPrKey string, RPuKey string, ipfsConn net.Conn) string {
	it := time.Now().UnixNano()

	content := GetDateContent(dateID)
	if content == "" {
		return ""
	}

	encry := EncryMes(content, balanceKey, SPrKey, RPuKey)
	addr := UploadMes(encry, ipfsConn)

	it = time.Now().UnixNano() - it
	ipfsUpTime = it
	ipfsUpNum += 1
	upDataSize = int64(len(encry))
	fmt.Println(encry)
	return addr
}

func Process(conn net.Conn, ipfsConn net.Conn)  {
	defer conn.Close()  //关闭连接
	/* for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])    //读取数据
		if err != nil{
			fmt.Println("客户端连接中断，进程中止")
			return
		}
		recvStr := string(buf[:n])
		// fmt.Println("收到客户端信息：",recvStr)
		conn.Write([]byte(recvStr)) //发送数据
	}*/
	var payee *tx.MicroPayee
	var PayerKey *address.PublicKey
	var bond *tx.Tx
	var refund *tx.Tx

	gKeyTime = time.Now().UnixNano()
	// 用于对称加密的密钥
	balanceKey := secretbox.GenerateSecretKey()
	// 用于非对称加密的公私钥
	encryPub, encryPri, encryErr := box.GenerateKeyPair()
	if (encryErr != nil) {
		panic(encryErr)
	}
	gKeyTime = time.Now().UnixNano() - gKeyTime
	
	var rec string
	rec = fmt.Sprintf("%v\n", gKeyTime)
	sRecord(rec)

	for {
		reader := bufio.NewReader(conn)
                buf := [40096]byte{}
		for {
			buf = [40096]byte{}
			n, err := reader.Read(buf[:]) //读取数据
			if err != nil {
				fmt.Println("客户端连接中断，进程中止")
				return
			}

			key := int(buf[0])
			recvMes := buf[1:n]
			// fmt.Println("收到客户端信息：", recvMes, "len :", n)

			rKey := []byte{1}
			rMes := []byte("")

			switch (key) {
			case 0: // close the connect
				return;
			case 1:
				temp, err := address.FromWIF(string(recvMes[:]), address.BitcoinTest)
				if err != nil {
					panic(err)
				}
				PayerKey = temp.PublicKey

				payee = tx.NewMicroPayee(PayerKey, serverKey, 200*tx.Unit, 0.001*tx.Unit)
				rKey[0] = 1
				rMes = []byte(wif2)
				break
			case 2:
				r := Byte2Tx(recvMes)
				refund = &r
				// fmt.Println("refund: ", refund.TxIn[0].Hash)

				sign, err := payee.SignRefund(refund, refund.Locktime)
				if err != nil {
					break
				}
				var mes1 Mes1
				mes1.Sign = sign
				mes1.Transaction = Tx2Byte(*refund)

				rKey[0] = 2
				rMes,_ = json.Marshal(mes1)
				// fmt.Println(j)

				break
			case 3:
				var mes2 Mes2
				json.Unmarshal(recvMes, &mes2)
				btemp := Byte2Tx(mes2.T1)
				bond = &btemp
				rtemp := Byte2Tx(mes2.T2)
				refund = &rtemp
				// fmt.Println("bond script: ", bond.TxOut[0].Value)
				err := payee.CheckBond(refund, bond);
				if  err != nil {
					rKey[0] = 0
  					panic(err)
					break
				}
				rKey[0] = 3
				break
			case 4: // 加密数据的位置
				// T1:数据id T2:买家支付的价格 T3:买家的公钥
				var mes3 Mes3
				json.Unmarshal(recvMes, &mes3)

				recvDateID,_ := strconv.Atoi(string(mes3.T1))
				mesnum,_ := strconv.Atoi(string(mes3.T2))
				recvPub := string(mes3.T3)

				// 上传数据并获取地址
				loadAddr := UploadMesById(recvDateID, balanceKey, encryPri, recvPub, ipfsConn)

				// 生成签名
				coinnum := uint64(mesnum)
				var realnum uint64 = 0.001*tx.Unit*coinnum;
				var signIP []byte
				signIP, err = payee.CSignIncremented(loadAddr, realnum)
				if (err != nil) {
					fmt.Println("4 write error");
					return;
				}

				// 构造返回信息, T1:更新交易签名 T2:下载地址 T3:卖家公钥
				mes3.T1 = signIP
				mes3.T2 = []byte(loadAddr)
				mes3.T3 = []byte(encryPub)

				// 返回信息
				rKey[0] = 4
				rMes, err = json.Marshal(mes3)
				break
			case 5:
				// increment := Byte2Ctx(recvMes)

				rKey[0] = 5
				fmt.Println("Channel Update Success!")
				ServerPrint()
				break
			default:
				rKey[0] = 0
				rMes = []byte("")
				fmt.Print("Illegal Key!\n")
				break
			}
			// fmt.Println("server send: ", BytesCombine(rKey, rMes))
			// 发送返回数据
			conn.Write(BytesCombine(rKey, rMes))
		}
	}
}

func ServerInit() {
	// 创建可用的公私钥对
	//ms5repuZHtBrKRE93FdWqz8JEo6d8ikM3k
	txKey, err := address.FromWIF(wif2, address.BitcoinTest)
	if err != nil {
		panic(err)
	}
	serverKey = txKey

	serverData, err = os.Create("data/serverData")
        if err != nil {
                panic(err)
        }
}

/*

func main()  {
	// 进行初始化
	ServerInit()

	listen,err := net.Listen("tcp","127.0.0.1:8888")
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
		go process(conn)    //启动一个goroutine处理连接
	}
}
*/
