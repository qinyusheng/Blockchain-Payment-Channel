# 关于本项目的一些细节

## 关于编译运行

直接将main.c文件以directory的形式编译运行即可，go build 配置比较简单，出现问题可以问我

## 外部包的安装

需要安装的外部包我都将其放到src.zip中，一同上传上去了。

虽然好像也可以直接使用 git build 直接安装，但是不知道是不是因为我这边的配置问题，我没有办法正常地使用git build，因此还是决定将他上传上来

直接将这件包，拷贝导 $GOPATH/src 文件夹里，里面可能已经有了一些同名地文件夹，例如 github,com。此时可以再自行将包逐个按照本来的路径复制过去，就可以顺利完成编译了

## 我们写的包

`IPFS` 请将其放到`$GOPATH/src/`目录下

`net1` 请将其放到`$GOPATH/src/`目录下

`tx1` 请将其放到`$GOPATH/src/github.com/bitgoin/`目录下

*我取名的时候有点随意了，之后如果有修改会及时更新该说明的*

## 关于运行

进入test目录之后，执行如下指令：

```shell
python start.py
```

正常情况下，会输出如下信息：

```shell
root@blockchain-VirtualBox:/home/blockchain/GOPATH/net# python start.py 
*** Creating network
*** Adding controller
*** Adding hosts:
h0 h1 ipfs 
*** Adding switches:
s0 
*** Adding links:
(20.00Mbit 100ms delay) (20.00Mbit 100ms delay) (h0, s0) (20.00Mbit 100ms delay) (20.00Mbit 100ms delay) (h1, s0) (20.00Mbit 100ms delay) (20.00Mbit 100ms delay) (ipfs, s0) 
*** Configuring hosts
h0 (cfs -1/100000us) h1 (cfs -1/100000us) ipfs (cfs -1/100000us) 
*** Starting controller
c0 
*** Starting 1 switches
s0 ...(20.00Mbit 100ms delay) (20.00Mbit 100ms delay) (20.00Mbit 100ms delay) 
*** IPFS node starting on ipfs
./ipfs
*** Blockchain nodes starting on h0
./nodes --ipfs 10.0.0.3 
*** Blockchain nodec starting on h1
./nodec --peer 10.0.0.1 --ipfs 10.0.0.3
*** Starting CLI:
mininet> 
```

### 开始运行ipfs服务器

在终端中输入

```shell
mininet> xterm ipfs
```

就会跳出一个窗口，对应ipfs虚拟机，在窗口中输入如下指令

```shell
./ipfs
```

### 开始运行卖家程序

配置h0

```c
mininet> xterm h0
```

在跳出的窗口中输入

```shell
./nodes --ipfs 10.0.0.3 
```

### 开始运行买家程序

配置h1

```c
mininet> xterm h1
```

在跳出的窗口中输入

```shell
./nodec --peer 10.0.0.1 --ipfs 10.0.0.3
```

*备注，需要输入的指令在运行start.py之后都有输出，往后可以此为主（如果我懒得更新这里的话）*

*请务必不要在意上面的语法错误以及未显示的中文字符，这不重要！*

## node

Process to run in each host participating in the network.

**Arguments:**
- `-i`  
IP address of the node in the network
- `-p`  
port to listen to peers (defaults to 9000)
- `--peer`  
IP address of peer
- `--ipfs`  
IP address of IPFS server