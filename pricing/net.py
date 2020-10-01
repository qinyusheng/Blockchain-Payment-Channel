import pay;
import random

class Network:
    graph = []
    node_num = 0
    path = []
    price = []
    imbalance = 0

    def __init__(self, num):
        self.graph = [[-1 for i in range(num)] for i in range(num)]
        self.price = [10 for _ in range(num)]
        self.node_num = num    

    # 默认两边抵押金相同
    def build(self, i, j, n):
        self.graph[i][j] = n
        self.graph[j][i] = n
    
    def build1(self, i, j, n1, n2):
        self.graph[i][j] = n1
        self.graph[j][i] = n2

    def trans(self, i, j, num):
        # print(i,j,len(self.graph))
        if (self.graph[i][j] < num):
            return False
        self.graph[i][j] -= num
        self.graph[j][i] += num

        if (self.graph[i][j] == 0):
            self.imbalance += 1
        return True

    def execute(self, way, wei):
        for i in range(len(way)-1):
            self.trans(way[i], way[i+1], wei)

    def default(self):
        self.node_num = 10;
        self.price = [10, 12, 5, 9, 13, 14, 10, 15, 10, 8]
        # 使用固定网络则注释
        # self.price = [random.randint(5,15) for _ in range(self.node_num)]
        self.graph = [[0 for i in range(self.node_num)] for i in range(self.node_num)]

        self.build( 0 , 1 , 8 )
        self.build( 0 , 2 , 7 )
        self.build( 0 , 3 , 7 )
        self.build( 0 , 8 , 9 )
        self.build( 1 , 2 , 5 )
        self.build( 1 , 5 , 7 )
        self.build( 1 , 7 , 9 )
        self.build( 1 , 8 , 7 )
        self.build( 2 , 1 , 9 )
        self.build( 2 , 4 , 6 )
        self.build( 2 , 6 , 9 )
        self.build( 2 , 7 , 8 )
        self.build( 3 , 2 , 5 )
        self.build( 3 , 4 , 8 )
        self.build( 3 , 5 , 6 )
        self.build( 3 , 6 , 9 )
        self.build( 3 , 9 , 8 )
        self.build( 4 , 0 , 9 )
        self.build( 4 , 7 , 9 )
        self.build( 4 , 9 , 9 )
        self.build( 5 , 0 , 6 )
        self.build( 5 , 1 , 8 )
        self.build( 5 , 4 , 5 )
        self.build( 5 , 6 , 8 )
        self.build( 5 , 8 , 9 )
        self.build( 6 , 1 , 9 )
        self.build( 6 , 2 , 6 )
        self.build( 6 , 3 , 8 )
        self.build( 6 , 5 , 6 )
        self.build( 6 , 7 , 5 )
        self.build( 7 , 1 , 6 )
        self.build( 7 , 2 , 7 )
        self.build( 7 , 5 , 5 )
        self.build( 7 , 8 , 7 )
        self.build( 8 , 1 , 5 )
        self.build( 8 , 3 , 9 )
        self.build( 8 , 5 , 5 )
        self.build( 8 , 6 , 8 )
        self.build( 9 , 0 , 6 )
        self.build( 9 , 1 , 5 )
        self.build( 9 , 3 , 7 )
        self.build( 9 , 4 , 7 )
        self.build( 9 , 6 , 8 )
        self.build( 9 , 7 , 6 )
        self.build( 9 , 8 , 9 )
        # 待补充

    def random(self, num):
        self.node_num = num;
        self.price = [10 for _ in range(self.node_num)]
        self.graph = [[-1 for i in range(self.node_num)] for i in range(self.node_num)]

        r = round(num/5)
        for i in range(num):
            for j in range(num):
                if (i != j and random.randint(0,max(r,1))==0):
                    # self.build(i, j, random.randint(5,9))
                    self.build(i, j, 5)

    def showGraph(self):
        print('x', end='\t')
        for i in range(self.node_num):
            print(i, end='\t')
        print()

        for i in range(self.node_num):
            print(i, end='\t')
            for j in range(self.node_num):
                print(self.graph[i][j], end='\t')
            print()
        return
    
    def get_a_to_b(self, s, r):
        return self.graph[s][r]        

class Path:
    m_net = Network(0)
    path = [] # 路径以字符数组的形式存储

    def __init__(self, net):
        self.m_net = net
        self.buildPath(net)
    
    # 创建对应一个network的path
    def buildPath(self, net):
        self.m_net = net
        for i in range(self.m_net.node_num):
            pset = []
            for j in range(self.m_net.node_num):
                pset.append(self.findPath(i,j))
            self.path.append(pset.copy())
        
    def findPath(self, i, j):
        ps = []
        p = [i]
        self.generate(j, p, ps)
        # print(i,j,ps)
        return ps

    def generate(self, j, p, ps):
        # if path is too long
        if (len(p)>3):
            return
        # if find node j
        now = p[len(p)-1]
        if (now == j):
            ps.append(p.copy())
            return
        
        # find all the possible channel
        for a in range(self.m_net.node_num):
            if (self.m_net.graph[now][a] > 0):
                if not (a in p):
                    p.append(a)
                    self.generate(j, p, ps)
                    p.pop()
        return
    
    def execute(self, way, wei):
        self.m_net.execute(way, wei)

'''
# 测试程序     
n = Network(2)
n.default()
p = Path(n)
print(p.findPath(0,1))
'''