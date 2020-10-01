# 测试文件运行方法

输入指令：

```shell
python start.py
```

## 相关文件说明

`start.py` 调用`mininet` python接口的python程序，主要负责仿真网络的创建

`ipfs` IPFS服务器，主要提供了建立连接，上传数据，下载数据等接口

`nodes` 卖家程序，连接`ipfs`服务器，初始化售卖数据，根据买家的目标数据id生成相应数据信封并上传ipfs

`nodec` 买家程序，主动跟卖家建立支付通道，然后再更新通道状态的过程中购买数据（以`ipfs`下载地址的形式）

后三个程序都由对应go源文件编译而来，编译指令（需要将各种包先安装好）

```shell
go build xxx.go
```

