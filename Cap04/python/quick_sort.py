def qsort(array):
    if len(array) < 2:
        return array
    else:
        pivot = array[0]
        less = [i for i in array[1:] if i <= pivot]
        higher = [i for i in array[1:] if i > pivot]
        
        return qsort(less) + [pivot] + qsort(higher)

array = [3, 6, 8, 10, 1, 2, 1]
print(qsort(array)) # [1, 1, 2, 3, 6, 8, 10]