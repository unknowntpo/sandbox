import numpy as np
import matplotlib.pyplot as plt

# load txt file as numpy array (type of data is 'numpy.ndarray')
data = np.loadtxt('my_data.txt')
# plot the curve using x axis and y axis
plt.plot(data[:, 0], data[:, 1])
plt.show()
