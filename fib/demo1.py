def normal_recursion(n):
    if n < 2:
        return n
    return normal_recursion(n - 1) + normal_recursion(n - 2)

if __name__ == '__main__':
    print(normal_recursion(100000))
