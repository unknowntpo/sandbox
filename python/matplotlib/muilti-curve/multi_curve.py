# Plot sin and cos curve on same picture
import numpy as np
import matplotlib.pyplot as plt

X = np.linspace(0, 2 * np.pi, 100)
Ya = np.sin(X)
Yb = np.cos(X)

# Plot two curves
plt.plot(X, Ya)
plt.plot(X, Yb)

plt.show()
