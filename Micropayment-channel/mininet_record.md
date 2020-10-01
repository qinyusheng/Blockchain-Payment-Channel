## 安装环境

1. 清华镜像下载ubuntu (我的是16.04.6)

   https://mirrors.tuna.tsinghua.edu.cn/ubuntu-releases/

2. VirtualBox共享文件夹

   https://jingyan.baidu.com/article/656db918cca831e381249cce.html

3. 在虚拟机中使用主机的 shadowsocks

   https://zerol.me/2018/06/12/vm-ss/

4. ubuntu apt-get 替换清华镜像

   https://blog.csdn.net/zgljl2012/article/details/79065174

   https://mirrors.tuna.tsinghua.edu.cn/help/ubuntu/

5. 安装Mininet

   http://mininet.org/download/

```bash
git clone git://github.com/mininet/mininet
mininet/util/install.sh -a
```

# Background

1. Introduction to Mininet

   https://github.com/mininet/mininet/wiki/Introduction-to-Mininet

2. Mininet Python API

   http://mininet.org/api/annotated.html

3. mininet实验 设置带宽之简单性能测试

   https://www.cnblogs.com/031602523liu/p/8993218.html

## Topology

switch和switch之间是否可以连接？

## Private Directory

默认状态下，所有host共享server的file system，但是这会导致host没有自己的系统配置（如网络配置）

如果需要设置每个host一个private directories，可以输入以下内容

```python
h = Host( 'h1', privateDirs=[ '/some/directory' ] )
```

文件夹分两类：

1. persistent directory（前面那个和后面那个有什么关系？）

```python
[ ( '/var/run', '/tmp/%(name)s/var/run' ) ]
```

2. temporary directory

```python
'/var/log'
```

一般两者都有

example: https://github.com/mininet/mininet/blob/master/examples/bind.py

## Host Configuration Methods

1. `IP()`: Return IP address of a host or specific interface.
2. `MAC()`: Return MAC address of a host or specific interface.
3. `setARP()`: Add a static ARP entry to a host's ARP cache.
4. `setIP()`: Set the IP address for a host or specific interface.
5. `setMAC()`: Set the MAC address for a host or specific interface

默认IP是什么，IP应该怎么设置

## CLI

Starting up the CLI can be useful for debugging your network, as it allows you to view the network topology (with the `net` command), test connectivity (with the `pingall` command), and send commands to individual hosts.

```bash
*** Starting CLI:
mininet> net
c0
s1 lo:  s1-eth1:h1-eth0 s1-eth2:h2-eth0
h1 h1-eth0:s1-eth1
h2 h2-eth0:s1-eth2
mininet> pingall
*** Ping: testing ping reachability
h1 -> h2
h2 -> h1
*** Results: 0% dropped (0/2 lost)
mininet> h1 ip link show
746: lo: <LOOPBACK,UP,LOWER_UP> mtu 16436 qdisc noqueue state UNKNOWN
	link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
749: h1-eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
	link/ether d6:13:2d:6f:98:95 brd ff:ff:ff:ff:ff:ff
```

## Mininet Cluster

Mininet Cluster can create a topology with nodes on remote machines.

（如何实现？）





