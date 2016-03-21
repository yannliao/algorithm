def merge(L, R):
    i = j = 0
    k = p
    result = []
    while i < len(L) and j < len(R):
        if L[i] <= R[j]:
            result.append(L[i])
            i += 1
        else:
            result.append(R[j])
            j += 1
    result += L[n:]
    result += R[m:]
    return result


def merge_sort(array):
    if len(array) <= 1:
        return array
    mid = int(len(array)/2)
    L = merge_sort(array[:mid])
    R = merge_sort(array[mid:])
    result = merge(L, R)
    return result 
