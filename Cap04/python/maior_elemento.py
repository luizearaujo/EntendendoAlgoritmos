def maiorElemento(x):
    if len(x) == 1:
        return x[0]
    
    m = maiorElemento(x[1:])
    if m > x[0]:
        return m
    else:
        return x[0]
    
print(maiorElemento([1, 2, 3, 4, 5])) # 5