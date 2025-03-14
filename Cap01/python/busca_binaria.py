def busca_binaria(lista, item):

    baixo = 0
    alto = len(lista) - 1
    meio = (baixo + alto) // 2
    
    while  baixo <= alto:
        meio = (baixo + alto) // 2
        chute = lista[meio]

        if chute == item:
            return chute
        if chute > item:
            alto = meio + 1
        else:
            baixo = meio +1

    return None
    

lista = [1 ,2 ,3, 4, 5, 6, 7, 8, 9, 10]
print(busca_binaria(lista, 2))