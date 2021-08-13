import numpy as np

with open("out.txt", encoding = 'utf-8') as f:
    x, y = np.loadtxt(f)

print(x, y)


