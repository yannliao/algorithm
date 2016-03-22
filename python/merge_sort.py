"""
  Merge Sort
  CLRS Chapter 2.3
"""

def merge(L, R):
    """
    Merge two sorted subsequences into one sorted tree
    Param L: the first sorted list
    Param R: the second sorted list
    Retrun result: the merged list of L and R which is sorted
    """
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
    """
    Sort the array in ascending order, then a sorted list is returned.
    Param array: sequences will be sorted.
    Return result: the sorted list.
    """
    if len(array) <= 1:
        return array
    mid = int(len(array)/2)
    L = merge_sort(array[:mid])
    R = merge_sort(array[mid:])
    result = merge(L, R)
    return result 
