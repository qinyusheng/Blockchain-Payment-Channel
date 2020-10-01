package net1

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bitgoin/address"
	"net"
	"time"
	"github.com/bitgoin/tx1"
	"strconv"
	"IPFS/client/clientProcess"
	"digital-envelope/common/box"
	"digital-envelope/common/envelope"
	"os"
)

// time
var cgKeyTime int64 = 0
var ctxProcTime int64 = 0
var ctxProcNum int64 = 0
var downloadTime int64 = 0
var downloadNum int64 = 0
var clientData *os.File

func cRecord(rec string) {
	_, err := clientData.WriteString(rec);
	if (err != nil) {
		panic(err)
	}
}

func ClientPrint() {
	fmt.Println("The time to generate key pair:", cgKeyTime, "(nm)")
	fmt.Println("The number of update transaction:", ctxProcNum, "(nm)")

	fmt.Println("The number of download and decrypt data:", downloadNum, "(nm)")

	var rec string
	rec = fmt.Sprintf("%v %v\n", ctxProcTime, downloadTime)
	cRecord(rec)
}

var IPaddress string = "127.0.0.1:8888"
var clientKey *address.PrivateKey
var serverPKey *address.PublicKey
var wif string = "928Qr9J5oAC6AYieWJ3fG3dZDjuC7BFVUqgu4GsvRVpoXiTaJJf"
var payer *tx.MicroPayer
var EncryPub string // 非对称加密公钥
var EncryPri string // 非对称加密私钥

func DownLoad(ipfs net.Conn, addr string, SPub string) string {
	t := time.Now().UnixNano()

	var mes2 Mes2
	bmes,_ := clientProcess.DownLoadByAddr(ipfs, addr)
	json.Unmarshal([]byte(bmes), &mes2)

	date,_ := envelope.Open(string(mes2.T1), string(mes2.T2), SPub, EncryPri)

	t = time.Now().UnixNano() - t
	downloadTime = t
	downloadNum += 1
	ClientPrint()
	return date
}

func CreateBond(conn net.Conn) {
	pt := time.Now().UnixNano()

	var recvKey int
	var recvMes []byte
	var sKey []byte
	var sMes []byte

	buf := [40096]byte{}

	sKey = []byte{1}
	sMes = []byte(wif)

	conn.Write(BytesCombine(sKey, sMes))
	buf = [40096]byte{}
	n,err := conn.Read(buf[:])
	if (err != nil) {
		fmt.Println("Read err: ", err)
		return
	}
	recvKey = int(buf[0])
	recvMes = buf[1:n]
	if recvKey != 1 {
		fmt.Println("Create fail!")
		return
	}
	// 获取服务端公钥
	Swif := string(recvMes[:])
	temp, err := address.FromWIF(Swif, address.BitcoinTest)
	if err != nil {
		panic(err)
	}
	serverPKey = temp.PublicKey

	// 创建通道用户
	payer = tx.NewMicroPayer(clientKey, serverPKey, 200*tx.Unit, 0.001*tx.Unit)
	payee := tx.NewMicroPayee(clientKey.PublicKey, temp, 200*tx.Unit, 0.001*tx.Unit)
	// 通道前期准备 key=2
	locktime := uint32(time.Now().Add(time.Hour).Unix())
	txhashes := []string{
		"12c2f61d839b2b38146715e4dfc0fd914906253920480298816f108513e53e5c",
		"12c2f61d839b2b38146715e4dfc0fd988806253920480298816f108513e53e5c",
	}
	values := []uint64{100 * tx.Unit, 150 * tx.Unit}
	script, err := hex.DecodeString("76a914d94987ba89c258372030bc9d610f89547757896488ac")
	if err != nil {
		panic(err)
	}
	utxos := make(tx.UTXOs, len(txhashes))
	for i, h := range txhashes {
		var ha []byte
		ha, err = hex.DecodeString(h)
		if err != nil {
			panic(err)
		}
		ha = tx.Reverse(ha)
		utxos[i] = &tx.UTXO{
			Key:     clientKey,
			TxHash:  ha,
			Value:   values[i],
			Script:  script,
			TxIndex: uint32(i + 1),
		}
	}
	// 创建待签名交易 key=2
	bond, refund, err := payer.CreateBond(locktime, utxos, clientKey.PublicKey.Address())
	sKey[0] = 2
	sMes = Tx2Byte(*refund)
	// fmt.Println("refund: ", refund.TxIn[0].Hash)
	if (err != nil) {
		fmt.Println("Read err: ", err)
		return
	}

	conn.Write(BytesCombine(sKey, sMes))
	n,err = conn.Read(buf[:])
	if (err != nil) {
		fmt.Println("Read err: ", err)
		return
	}
	recvKey = int(buf[0])
	recvMes = buf[1:n]
	if recvKey != 2 {
		fmt.Println("Create fail!")
		return
	}
	var mes1 Mes1
	json.Unmarshal(recvMes, &mes1)
	json.Unmarshal(recvMes, &mes1)
	*refund = Byte2Tx(mes1.Transaction)

	// 进行第三步
	if err = payer.SignRefund(refund, mes1.Sign); err != nil {
		fmt.Println("Error Refund Sign!")
		return
	}
        // fmt.Println("success 2")

if err = payee.CheckBond(refund, bond); err != nil {
		panic(err)
	}

	sKey[0] = 3
	var mes2 Mes2
	mes2.T1 = Tx2Byte(*bond)
	mes2.T2 = Tx2Byte(*refund)
	sMes,err = json.Marshal(mes2)
	conn.Write(BytesCombine(sKey, sMes))
	// fmt.Println("bond origin: ", bond.TxIn[0].Index)

       // json.Unmarshal(recvMes, &mes2)
// fmt.Println("bond json: ", btemp.TxIn[0].Index)

	buf = [40096]byte{}
	n,err = conn.Read(buf[:])
	if (err != nil) {
		fmt.Println("Read err: ", err)
		return
	}
	recvKey = int(buf[0])
	recvMes = buf[1:n]
	if recvKey != 3 {
		fmt.Println("Create fail!")
		return
	}

	pt = time.Now().UnixNano() - pt
	var rec string
	rec = fmt.Sprintf("%v\n", pt)
	cRecord(rec)

	// success
	fmt.Println("Channel Create Success!")
	return
}

func BuyData(conn net.Conn, dateID int, fee uint64) (string, string){
	pt := time.Now().UnixNano()

	var recvKey int
	var recvMes []byte
	var sKey []byte
	var sMes []byte
	var dateAddr string
	var realnum uint64 = 0.001*tx.Unit*fee;

	buf := [1024]byte{}

	// 发送信息
	var mes3 Mes3
	mes3.T1 = []byte(strconv.Itoa(dateID))
	mes3.T2 = []byte(strconv.Itoa(int(fee)))
	mes3.T3 = []byte(EncryPub)

	sKey = []byte{4}
	sMes,_ = json.Marshal(mes3)
	conn.Write(BytesCombine(sKey, sMes))

	n,err := conn.Read(buf[:])
	if (err != nil) {
		fmt.Println("Read err: ", err)
		return "",""
	}
	recvKey = int(buf[0])
	recvMes = buf[1:n]
	if recvKey != 4 {
		fmt.Println("Buy Fail!")
		return "",""
	}
	json.Unmarshal(recvMes, &mes3);
	ServerSign := mes3.T1
	dateAddr = string(mes3.T2)
	SendPub := string(mes3.T3)

	if dateAddr == "" || SendPub == "" {
		return "",""
	}

	// 发送自己构建出来的交易
	increment, err := payer.CIncrementedTx(realnum, ServerSign, []string{dateAddr, dateAddr})
	if err != nil {
		panic(err);
	}
	sKey[0] = byte(5)
	sMes = Ctx2Byte(*increment)
	conn.Write(BytesCombine(sKey, sMes))

	n,err = conn.Read(buf[:])
	if (err != nil) {
		fmt.Println("Read err: ", err)
		return "fail 1",""
	}
	recvKey = int(buf[0])
	recvMes = buf[1:n]
	if recvKey != 5 {
		fmt.Println("Illegal transaction!")
		return "fail 2",""
	}

	// fmt.Println(dateAddr);
	pt = time.Now().UnixNano() - pt
	ctxProcNum += 1
	ctxProcTime = pt

	// ClientPrint()
	return dateAddr, SendPub
}

func ClientInit() {
	//n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy
	var err error
	clientKey, err = address.FromWIF(wif, address.BitcoinTest)
	if err != nil {
		panic(err)
	}

	cgKeyTime = time.Now().UnixNano()
	EncryPub, EncryPri, err = box.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	cgKeyTime = time.Now().UnixNano() - cgKeyTime

	clientData, err = os.Create("data/clientData")
	if err != nil {
		panic(err)
	}

	var rec string
	rec = fmt.Sprintf("%v\n", cgKeyTime)
	cRecord(rec)
}

func ClientClose() {
	clientData.Close()
}
