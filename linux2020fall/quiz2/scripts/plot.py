import matplotlib.pyplot as plt
import numpy as np

data = np.loadtxt("../plot/out.txt")

# Set limit of axis
plt.xlim(2000, 3000)
plt.ylim(0, 100000)
plt.plot(data[:, 0], data[:, 1], label="8 bytes/ time")
plt.plot(data[:, 0], data[:, 2], label="1 byte/ time")
plt.legend()
plt.title("isascii calculate speed comparison: ")
plt.show()
