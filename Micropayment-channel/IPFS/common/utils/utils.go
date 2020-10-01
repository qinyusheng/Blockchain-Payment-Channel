package utils

import (
	"IPFS/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//这里将这些方法关联到结构体
type Transfer struct{
	//
	Conn net.Conn
	Buf [8096]byte  //传输时使用的缓冲
}

//读包，解包函数
func (this *Transfer)ReadPkg() (mes message.Message, err error){
	buf := make([]byte, 8086)
	_, err = this.Conn.Read(buf[:4]) //先读出包的长度
	if err != nil{
		return
	}
	//根据buf[:4] 转成一个 unit32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据 长度 pkgLen 读取下一个内容包
	n, err := this.Conn.Read(buf[:pkgLen])
	if n!=int(pkgLen) || err != nil{
		return
	}
	//将 buf[:pkgLen] 反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("readPkg json.Unmarshal err=", err)
		return
	}
	return

}

//读包，解包函数
func (this *Transfer)ReadPkg2() (mes message.Message, err error){
	var buf_size uint32 = 8086
	buf := make([]byte, buf_size)
	_, err = this.Conn.Read(buf[:4]) //先读出包的长度
	if err != nil{
		return
	}
	//根据buf[:4] 转成一个 unit32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
fmt.Println("len of pkg to read:", pkgLen)

	//根据 长度 pkgLen 读取下一个内容包
	var n int
	res := make([]byte, pkgLen+1)
	var res_len uint32 = 0

	for true {
		if (pkgLen-res_len) > buf_size {
			n, err = this.Conn.Read(buf[:buf_size])
			if n != int(buf_size) || err != nil {
				panic("Read to small1")
				return
			}

			for i:=0; i<int(buf_size); i ++ {
				res[int(res_len)+i] = buf[i]
			}
			res_len += buf_size
 		} else {
 			var l = int(pkgLen-res_len)
			n, err = this.Conn.Read(buf[:l])
			if n != l || err != nil {
				panic("Read to small2")
				return
			}
			for i:=0; i<l; i ++ {
				res[int(res_len)+i] = buf[i]
			}
			break
		}
	}
	
	//将 buf[:pkgLen] 反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("readPkg json.Unmarshal err=", err)
		return
	}
	return

}

func (this *Transfer) WritePkg(mes message.Message)(err error)  {
	data, err := json.Marshal(mes)
	if err != nil{
		fmt.Println("mes json.Marshal err=", err)
		return
	}
	//7. 此时data为待发送数据包
	//7.1 现将data的字节数发送给对方进行检错
	//先将 data长度->转成一个byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := this.Conn.Write(buf[:4])  //发送数据字节长度
	if n!=4 || err != nil {
		fmt.Println("conn.Write(head) err=", err)
		return
	}
	_, err = this.Conn.Write(data)  //发送数据字节长度
	if err != nil {
		fmt.Println("conn.Write(body) err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg2(mes message.Message)(err error)  {
	var buf_size uint32 = 8086

	data, err := json.Marshal(mes)
	if err != nil{
		fmt.Println("mes json.Marshal err=", err)
		return
	}
	//7. 此时data为待发送数据包
	//7.1 现将data的字节数发送给对方进行检错
	//先将 data长度->转成一个byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := this.Conn.Write(buf[:4])  //发送数据字节长度
	if n!=4 || err != nil {
		fmt.Println("conn.Write(head) err=", err)
		return
	}
fmt.Println("pkg len", pkgLen)
	var start uint32 = 0
	for true {
		if pkgLen-start > buf_size {
			n, err = this.Conn.Write(data[start : start+buf_size]) //发送数据字节长度
			if n != int(buf_size) {
				panic("write too small")
			}
			if err != nil {
				fmt.Println("conn.Write(body) err=", err)
				return
			}
			start += buf_size
		} else {
			n, err = this.Conn.Write(data[start:pkgLen]) //发送数据字节长度
			if n != int(pkgLen-start) {
				panic("write too small")
			}
			if err != nil {
				fmt.Println("conn.Write(body) err=", err)
				return
			}
			break
		}
	}

	return
}