package serverProcess

import (
	"IPFS/common/message"
	"IPFS/common/utils"
	"fmt"
	"io"
	"net"
)



type Processor struct{
	Conn net.Conn    //处理和客户的连接
	userId int       //判断该客户身份
}

//ServerProcessMes函数 功能：根据客户端发送消息种类不同，决定调用那个函数来处理
func (this *Processor)ServerProcessMes(mes *message.Message) (err error){

	switch mes.Type {
	case message.LoginMesType:
		//处理登录请求
		//创建一个UserProcess实例
		up := &ReqProcess{
			Conn: this.Conn,
		}
		this.userId, err = up.ServerProcessLogin(mes)
		if err != nil{
			fmt.Println("serverProcessLogin err=", err)
			return
		}
	case message.UpLoadMesType:
		up := &ReqProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessUpLoad(mes, this.userId)
		if err != nil{
			fmt.Println("serverProcessUpload err=", err)
			return
		}
	case message.DownloadReqType:
		up := &ReqProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessDlReq(this.userId)
		if err != nil{
			fmt.Println("ServerProcessDlReq err=", err)
			return
		}
	case message.DownloadAddrType:
		up := &ReqProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessDlAddr(mes, this.userId)
		if err != nil{
			fmt.Println("ServerProcessDlAddr err=", err)
			return
		}

	default:
		fmt.Println("该类型消息不存在")
	}
	return
}
//总控函数
func (this *Processor)ServerProcessor() (err error) {
	//读客户端的消息
	for{
		//创建一个transfer实例，完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出")
				return err
			}else {
				fmt.Println("process readPkg err=",err)
				return err
			}
		}
		//fmt.Println("mes= ", mes)
		err = this.ServerProcessMes(&mes)
		if err != nil{
			return err
		}
	}
}