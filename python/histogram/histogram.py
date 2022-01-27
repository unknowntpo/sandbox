import random
import math

def prepare_list():
    l = []
    for i in range(10):
        l.append(random.randint(0, 100))
    return l 
def get_mean(l):
    u = sum(l) / len(l)
    return u

def get_variance(l: list, u: float):
    # Calculate variance
    acc = 0
    for i, v in enumerate(l):
        # (v - u)^2 
        tmp = abs(v - u)
        acc += tmp*tmp
    variance = acc / len(l)
    return variance

def get_standard_deviation(variance: float):
    return math.sqrt(variance)

"""
# ./bitehist.py
Tracing... Hit Ctrl-C to end.
^C
     kbytes          : count     distribution
       0 -> 1        : 3        |                                      |
       2 -> 3        : 0        |                                      |
       4 -> 7        : 211      |**********                            |
       8 -> 15       : 0        |                                      |
      16 -> 31       : 0        |                                      |
      32 -> 63       : 0        |                                      |
      64 -> 127      : 1        |                                      |
     128 -> 255      : 800      |**************************************|
"""
l = prepare_list()
print("list:", l)
u = get_mean(l)
print("mean value: ", u)

variance = get_variance(l, u)
print("variance: ", variance)

sd = get_standard_deviation(variance)
print("standard deviation: ", sd)
