s = "hello"

l = list(s)
i, j = 0, len(s) - 1
while (i < j):
    l[i], l[j] = l[j], l[i]
    i+=1
    j-=1

print("".join(l))

s = "".join(l)

print(s)
"""
i
0 1 2 3 4 len(hello)
-5 -4 -3 -2 -1
            j
h e l l o
"""
