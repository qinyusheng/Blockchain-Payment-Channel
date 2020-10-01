import random
import numpy as np

pay_matrix = [[45, 59, 35, 24, 19, 37, 38, 24, 35, 39], [44, 54, 22, 42, 33, 32, 46, 25, 31, 30], [28, 54, 31, 29, 36, 37, 19, 26, 24, 19], [27, 43, 38, 44, 18, 18, 18, 39, 32, 6], [7, 32, 32, 50, 27, 18, 44, 22, 30, 26], [18, 23, 52, 36, 37, 15, 27, 32, 33, 29], [31, 22, 40, 20, 27, 38, 19, 19, 32, 23], [31, 38, 22, 27, 34, 46, 35, 32, 47, 3], [29, 26, 38, 25, 54, 17, 33, 29, 17, 18], [16, 42, 42, 47, 33, 24, 35, 32, 36, 47]]
pay_matrix_sum = 3116

class Payment:
    sender = 0;
    receiver = 0;
    wei = 0;

    def __init__(self):
        self.sender = 0
        self.receiver = 0
        self.wei = 10
    
    def __init__(self, s, r, w):
        self.sender = s
        self.receiver = r
        self.wei = w

    def show(self):
        print("s:", self.sender, "\tr:", self.receiver, "\tw:", self.wei)

def aPay(node_num):
    s = random.randint(0, node_num-1)
    r = random.randint(0, node_num-1)
    while r==s:
        r = random.randint(0, node_num-1)
    return Payment(s, r, 1)

# 随机分布
def generate(node_num):
    global pay_matrix
    global pay_matrix_sum

    pay_matrix = [[random.randint(1,100) for i in range(node_num)] for i in range(node_num)]
    for i in range(node_num):
        for j in range(node_num):
            if (i == j):
                pay_matrix[i][j] = 0
            pay_matrix_sum += pay_matrix[i][j]

# 正态分布
def generate_normal(node_num):
    global pay_matrix
    global pay_matrix_sum

    # 生成正态分布 u=20 标准差=12
    m = np.random.normal(20, 12, size=(node_num, node_num))
    pay_matrix = [[0 for i in range(node_num)] for i in range(node_num)]
    
    for i in range(node_num):
        for j in range(node_num):
            if (i == j or m[i][j] < 0):
                pay_matrix[i][j] = 0
            else:
                pay_matrix[i][j] = int(m[i][j])
            pay_matrix_sum += pay_matrix[i][j]



def pay(node_num):
    global pay_matrix
    global pay_matrix_sum

    r = random.randint(1, pay_matrix_sum)
    a = 0
    b = 0

    for i in range(node_num):
        for j in range(node_num):
            r -= pay_matrix[i][j]
            if (r <= 0):
                a = i
                b = j
                break
        if (r <= 0):
            break
    
    # return Payment(a, b, 1)
    return Payment(a, b, random.randint(1,2))

'''
generate(5)
print(pay_matrix)
print(pay_matrix_sum)
for i in range(10):
    pay(5).show()
'''