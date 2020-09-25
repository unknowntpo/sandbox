import matplotlib.pyplot as plt

X = range(100) #  X: [0, 1, 2... 99]
Y = [value ** 2 for value in X]  # get each value of X[0] ... X[99] and **2 

print(X)

print(type(X[0]))
plt.plot(X, Y)
plt.show()

