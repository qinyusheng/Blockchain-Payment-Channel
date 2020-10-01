import net
import pay

# 实验一需求
channela = 1;
channelb = 2;
channelprice = 5;
froma = 0;
fromb = 0;

# 价格差
channeldiff = 2;

# 记录交易平衡度
a_to_b = 0;
b_to_a = 0;

def route(pay, path, kind):
    # 提取需要的信息
    wei = pay.wei
    all_ps = path.path[pay.sender][pay.receiver].copy()
    net = path.m_net
    ps = []

    for l in all_ps:
        if isable(l, wei, net):
            ps.append(l)
    # 如果没有路径直接结束
    if(len(ps)==0):
        return ([], 0)

    # 计算出最便宜的路径
    index = -1
    m = 99999
    for i in range(len(ps)):
        pr = getPrice(ps[i], wei, net, kind)
        if (pr<m):
            index = i
            m = pr
    
    # 实验一
    find_channel(ps[index], wei);

    return (ps[index], m)

def getPrice(way, wei, net, kind):
    all = (net.price[way[0]] + net.price[way[len(way)-1]]) / 2

    for i in range(len(way)-1):
        all += pricing(way[i], way[i+1], wei, net, kind)

    return all

def isable(way, wei, net):
    for i in range(len(way)-1):
        if(net.graph[way[i]][way[i+1]] < wei):
            return False 
    return True

# 实际上还是针对通道算钱
def pricing(s, r, wei, net, kind):
    global channelprice
    global channeldiff
    
    # 价格
    sp = net.price[s]
    rp = net.price[r]
    # 权重
    sw = net.graph[s][r]
    rw = net.graph[r][s]
    
    res = 0
    # k指每损失一个平衡状态时，定价的变化幅度
    k = 1
    # 实验一需求：注释掉
    '''
    if kind==0: # 直接取普通静态价格
        res = (sp+rp)/2 * wei
    elif kind==1: # 取与初始值的比例，倍数[0,2]
        res = (sp+rp)/2 * (2*rw/(sw+rw)) * wei
    elif kind==2: # 接收权重与发送权重的比例，倍数[0,无穷]
        res = (sp+rp)/2 * rw/sw * wei
    elif kind==3:
        res = (sp+rp)/2 * wei + 0.5*(rw-sw)
    elif kind==4:
        res = (sp+rp)/2*wei - k * (abs(sw - rw) - abs((sw - wei) - (rw + wei)))
    '''
    # 实验一需求，重新确定定价方案
    if (s == channela and r == channelb):
        return channelprice*wei * (2*rw/(sw+rw));
    elif (s == channelb and r == channela):
        return channelprice*wei * (2*rw/(sw+rw));
    else:
        return (sp+rp)/2 * wei * (2*rw/(sw+rw));

    if res > 0:
        return res
    else:
        return 0

def execute(pay, path, kind):
    res = route(pay, path, kind)
    if  len(res[0])==0:
        return (0, False)
    # 不考虑平衡问题则注释
    path.execute(res[0], pay.wei)
    return (res[1], True)

# 实验一需求
def find_channel(way, wei):
    global channela;
    global channelb;
    global channelprice; # 其他通道全是10
    global froma;
    global fromb;

    # 记录交易平衡度
    global a_to_b;
    global b_to_a;

    for i in range(len(way)-1):
        if(way[i] == channela and way[i+1] == channelb):
            froma += 1;
            a_to_b += wei;
        elif(way[i] == channelb and way[i+1] == channela):
            fromb += 1;
            b_to_a += wei;

'''
# 测试程序
tn = net.Network(1)
tn.default()
tpath = net.Path(tn)
tn.showGraph()

tpay = pay.Payment(0,1,1)
print("first:")
print(execute(tpay, tpath))
# tn.showGraph()

print("second:")
print(execute(tpay, tpath))
# tn.showGraph()

print("third:")
print(execute(tpay, tpath))
# tn.showGraph()

print("forth:")
print(execute(tpay, tpath))
# tn.showGraph()

print("fifth:")
print(execute(tpay, tpath))
tn.showGraph()
'''