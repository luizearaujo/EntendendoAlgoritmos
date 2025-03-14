def soma(x):
    if len(x) == 0:
        return 0
    if len(x) == 1:
        return x[0]
    else:
        return x[0] + soma(x[1:])
    
print(soma([1, 2, 3, 4, 5])) # 15