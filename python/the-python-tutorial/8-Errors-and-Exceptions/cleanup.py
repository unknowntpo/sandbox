def divide(x, y):
    try:
        result = x / y
    except ZeroDivisionError:
        print("error: divided by zero")
        return
    else:
        print("result is", result)
        return
    finally:
        print('execution finally clause')


divide(11, 2)
divide(12, 0)
