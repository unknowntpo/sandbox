import numpy as np
import matplotlib.pyplot as plt

X = np.linspace(0, 2 * np.pi, 100) # 100 points, from 0 to 2pi
Y = np.sin(X)
plt.plot(X, Y)
plt.show()
