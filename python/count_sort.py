# -*- coding: utf-8 -*-
"""
    Counting Sort
    Chapter 8
"""


def counting_sort(array, k):
    """
      Sorting array in theta n time
    """
    count = [0 for i in range(0, k)]
    bucket = {}
    # for i in range(0, k):
    #     C[i] = 0

    for j in range(0, len(array)):
        count[array[j]] += 1
    # print('count: {0}'.format(count))
    for i in range(1, k):
        count[i] = count[i] + count[i - 1]
    # print('count1: {0}'.format(count))
    for j in range(len(array) - 1, -1, -1):
        # bucket[count[array[j]]] = array[j]
        bucket[count[array[j]]] = array[j]
        # print('bucket is {0}'.format(bucket))
        count[array[j]] -= 1
    return bucket
