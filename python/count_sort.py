# -*- coding: utf-8 -*-
"""
    Counting Sort
    Chapter 8
"""


def counting_sort(sequence, k):
    """
      Sorting array in theta n time
    """
    count = [0 for i in range(0, k)]
    bucket = [] * len(sequence)

    for ele in sequence:
        count[ele] += 1

    for i in range(1, k):
        count[i] += count[i - 1]

    for j in range(len(sequence) - 1, -1, -1):
        item = sequence[j]
        bucket[count[item] - 1] = sequence[j]
        count[item] -= 1
    return bucket


def radix_sort(A, d):
    for i in range(d):
        counting_sort(A, 11)
