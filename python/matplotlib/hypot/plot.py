import numpy as np
import matplotlib.pyplot as plt


X = np.linspace(0, 100, 100)
Y1 = np.sqrt(X*X*X)
Y2 = np.sqrt(X*X)
plt.plot(X, Y1, label = 'X^3')
plt.plot(X, Y2, label = 'X^2')
plt.legend()
plt.show()
