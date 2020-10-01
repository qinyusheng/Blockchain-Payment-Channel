package message

const(
	LoginMesType		="LoginMes "
	ResMesType          ="ResMes"
	UpLoadMesType       ="UpLoadMes"
	DownloadReqType     ="DownloadReq"
	DownloadResType     ="DownloadRes"
	DownloadAddrType    ="DownloadAddr"
	DownloadContType    ="DownloadCont"
)

type Message struct{   //发送的主包，用于封装下面的消息类型
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息数据
}

type LoginMes struct{  //登录请求消息
	UserId int `json:"userId"` //用户id
}

type ResMes struct{ //服务器回送的登录结果
	Code int `json:"code"` //200 表示登录成功 201表示登陆失败 300表示上传成功 301表示上传失败
	Error string `json:"erro"` // 返回错误信息
}

type UpLoadMes struct {  //上传消息
	Cipher string  //密文
	//DigEnvelope string  //数字信封
}

type DownloadReq struct{  //客户端发出的消息请求

}

type DownloadRes struct{ //服务器返回的请求结果，一般会给出消息在数据库中的地址
	MesNum int //表示未读消息的数目
	ResMes []string
}

type DownloadAddr struct{  //客户端指出所需要的下载的消息
	Addr int
}

type DownloadCont struct {  //返回下载的结果
	Cipher string  //密文
	//DigEnvelope string  //数字信封
	Code int  //错误代码 400表示下载成功 404 表示下载失败
}




