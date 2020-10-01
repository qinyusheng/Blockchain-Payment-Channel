import pay
import net
import route
import copy
import random

# 新实验一
net1 = net.Network(0);
net1.default();

# 如果要使用默认网络则注释
# net1.random(10);

while (net1.get_a_to_b(route.channela, route.channelb) <= 0):
    net1.random(10);
print("a_to_b is", net1.get_a_to_b(route.channela, route.channelb))

path1 = net.Path(net1);
path2 = copy.deepcopy(path1);
# 注释掉可以以固定概率生成交易的
# pay.generate_normal(net1.node_num);
payment_set = [pay.pay(net1.node_num) for _ in range(10000)];

print("默认网络")
print("其他所有通道手续费固定")
print("交易金额随机 1-2")
print("考虑通道平衡性")
print("交易正态分布")

for k in range(0,20):
    route.channelprice = k;
    for l in payment_set:
        route.execute(l, path1, 0)
    print("price = ", route.channelprice)
    print(route.channela, "to", route.channelb, "=", route.froma, "all =", route.a_to_b)
    print(route.channelb, "to", route.channela, "=", route.fromb, "all =", route.b_to_a)
    print("sum =", route.froma+route.fromb)
    # 清零
    route.a_to_b = 0;
    route.b_to_a = 0;
    route.froma = 0;
    route.fromb = 0;
    path1 = copy.deepcopy(path2);


'''
net1 = net.Network(0)
net1.random(15)
path1 = net.Path(net1)
print("通道已成功计算")
path2 = copy.deepcopy(path1)
path3 = copy.deepcopy(path2)
path4 = copy.deepcopy(path3)
path5 = copy.deepcopy(path4)
print("网络已就绪")

pay.generate_normal(net1.node_num)
payment_set = [pay.pay(net1.node_num) for _ in range(10000)]
print("交易成功生成")
'''

'''
for i in payment_set:
    i.show()

'''

'''
# 实验 1
static = 0
dynamic1 = 0
dynamic2 = 0
dynamic3 = 0
dynamic4 = 0

# 路由金额
n1 = 0
n2 = 0
n3 = 0
n4 = 0
n5 = 0

nn = 0

for l in payment_set:
    if (nn % 10000 == 0):
        print(nn)
    nn += 1
    r1 = route.execute(l, path1, 0)
    r2 = route.execute(l, path2, 1)
    r3 = route.execute(l, path3, 2)
    r4 = route.execute(l, path4, 3)
    r5 = route.execute(l, path5, 4)
    if (r1[1]):
        static += 1
        n1 += r1[0]
    if (r2[1]):
        dynamic1 += 1
        n2 += r2[0]
    if (r3[1]):
        dynamic2 += 1
        n3 += r3[0]
    if (r4[1]):
        dynamic3 += 1
        n4 += r4[0]
    if (r5[1]):
        dynamic4 += 1
        n5 += r5[0]

print("static:", static, path1.m_net.imbalance)
print("dynamic1:", dynamic1, path2.m_net.imbalance)
print("dynamic2:", dynamic2, path3.m_net.imbalance)
print("dynamic3:", dynamic3, path4.m_net.imbalance)
print("dynamic4:", dynamic4, path5.m_net.imbalance)
'''


'''
# 实验 2
path1.m_net.price = [random.randint(5,10) for i in range(path1.m_net.node_num)]
path2.m_net.price = [random.randint(5,10) for i in range(path2.m_net.node_num)]

test1 = 0
test2 = 0

t = 0

min_num = 100000
max_num = 0

for l in payment_set:
    if (route.execute(l, path1, 0)[1]):
        test1 += 1
    if (route.execute(l, path2, 0)[1]):
        test2 += 1

for i in range(100):
    path1.m_net.default()
    path1.m_net.price = [random.randint(5,10) for i in range(path1.m_net.node_num)]
    for l in payment_set:
        if (route.execute(l, path1, 0)[1]):
            t += 1
    if (t > max_num):
        max_num = t
    if (t < min_num):
        min_num = t
    t = 0

print("min:", min_num)
print("max:", max_num)
'''