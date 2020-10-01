package serverProcess

import (
	"IPFS/common/message"
	"IPFS/common/utils"
	"IPFS/server/levdb"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

type ReqProcess struct{
	Conn net.Conn
}
//ServerProcessLogin函数 专门处理登录请求
func (this *ReqProcess)ServerProcessLogin(mes *message.Message) (UserId int, err error){
	//将mes中的mes.data反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.ResMesType
	var loginResMes message.ResMes


	if loginMes.UserId == 1 ||loginMes.UserId==2 {
		loginResMes.Code = 200
	} else{
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在!"
	}
	UserId = loginMes.UserId

	//将loginResMes 封装进resMes
	date, err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json.Marshal() err=", err)
		return
	}
	resMes.Data = string(date)
	//将resMes 发送给客户端

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(resMes)
	return
}

//ServerProcessUoLoad函数 专门处理上传请求
func (this *ReqProcess)ServerProcessUpLoad(mes *message.Message, userId int) (err error){
	var key int = 0;
	//将mes中的mes.data反序列化
	var UploadMes message.UpLoadMes
	err = json.Unmarshal([]byte(mes.Data), &UploadMes)
	if err != nil{
		fmt.Println("json.Unmarshal err=", err)
	}
	fmt.Printf("%d upLoadMes: %s\n",userId, UploadMes.Cipher)
	if userId==1{
		r := levdb.CountA2B()         //查询A发给B的消息有多少条
		key = r + 1       //以当前的消息数量加一作为key
		keys := strconv.Itoa(key)
		levdb.SaveA2B(keys, UploadMes.Cipher)
	}else{
		r := levdb.CountB2A()         //查询B发给A的消息有多少条
		key = r + 1       //以当前的消息数量加一作为key
		keys := strconv.Itoa(key)
		levdb.SaveB2A(keys, UploadMes.Cipher)
	}

	var resMes message.Message
	resMes.Type = message.ResMesType
	var UpResMes message.ResMes
	UpResMes.Code = 302+key
	//将ResMes 封装进resMes
	date, err := json.Marshal(UpResMes)
	if err != nil{
		fmt.Println("json.Marshal() err=", err)
	}
	resMes.Data = string(date)
	//将resMes 发送给客户端
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(resMes)
	if err != nil{
		fmt.Println("tf.WritePkg err=", err)
	}
	return
}

//ServerProcessDlReq函数 专门处理下载请求，并返回可取用的消息队列
func (this *ReqProcess)ServerProcessDlReq(userId int) (err error){

	var r []string
	if userId==1{
		r = levdb.SeekB2A()        //查询A的未读消息有多少条
		for _, v := range r{
			fmt.Printf("%s\n", v)
		}
	}else{
		r = levdb.SeekA2B()        //查询B的未读消息有多少条
		for _, v := range r{
			fmt.Printf("%s\n", v)
		}
	}
	num := len(r)
	var resMes message.Message
	resMes.Type = message.DownloadResType
	var DlResMes message.DownloadRes
	DlResMes.MesNum = num
	DlResMes.ResMes = r
	date, err := json.Marshal(DlResMes)
	if err != nil{
		fmt.Println("json.Marshal() err=", err)
	}
	resMes.Data = string(date)
	//将resMes 发送给客户端
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(resMes)
	if err != nil{
		fmt.Println("tf.WritePkg err=", err)
	}
	return
}

//ServerProcessDlAddr函数 根据客户端给出的消息地址返回消息内容
func (this *ReqProcess)ServerProcessDlAddr(mes *message.Message, userId int) (err error) {
	//将mes中的mes.data反序列化
	var DlAddrMes message.DownloadAddr
	err = json.Unmarshal([]byte(mes.Data), &DlAddrMes)
	if err != nil{
		fmt.Println("json.Unmarshal err=", err)
	}
	add := DlAddrMes.Addr
	adds := strconv.Itoa(add)
	var r string
	var code int

	if userId==1{
       //根据地址取出消息
		rr, err := levdb.GetB2A(adds)
		r = string(rr)
		if err != nil{
			fmt.Println("levdb.GetB2A err=", err)
			code = 404
		}else {
			code = 400
		}

	}else{
		rr, err := levdb.GetA2B(adds)
		r = string(rr)
		if err != nil{
			fmt.Println("levdb.GetB2A err=", err)
			code = 404
		}else{
			code = 400
		}
	}

	var resMes message.Message
	resMes.Type = message.DownloadContType
	var DlContMes message.DownloadCont
	DlContMes.Code = code
	DlContMes.Cipher = r
	date, err := json.Marshal(DlContMes)
	if err != nil{
		fmt.Println("json.Marshal() err=", err)
	}
	resMes.Data = string(date)
	//将resMes 发送给客户端
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(resMes)
	if err != nil{
		fmt.Println("tf.WritePkg err=", err)
	}
	return
}
