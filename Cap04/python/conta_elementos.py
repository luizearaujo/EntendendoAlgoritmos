def contaElementos(x):
    if len(x) == 1:
        return 1
    
    return 1 + contaElementos(x[1:])

print(contaElementos([1, 2, 3, 4, 5])) # 5