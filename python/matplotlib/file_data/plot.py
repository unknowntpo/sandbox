import matplotlib.pyplot as plt

"""
my_data.txt

0  0
1  1
2  4
4 16
5 25
6 36
"""
X, Y = [], []
for line in open('my_data.txt', 'r'):  # line[0]: 0 0 line[1] 1 1 ...
    values = [float(s) for s in line.split()]
    X.append(values[0])
    Y.append(values[1])
plt.plot(X, Y)
plt.show()
