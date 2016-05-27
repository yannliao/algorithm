# -*- coding: utf-8 -*-
"""
  Heap Sort
  CLRS Chapter 6
"""


def max_heapify(A, heap_size, i):
    """
    lets the value A[i] "float down" in the max-heap, so that the subtree at
    index i obeys the max-heap property.

    Param: A: a sequence of n numbers.
    Param: heap_size: heap size of sequence A.
    Param: i: root index of the subtree.
    """
    l = 2 * i + 1
    r = 2 * (i + 1)
    largest = i
    if l < heap_size and A[l] > A[i]:
        largest = l
    else:
        largest = i

    if r < heap_size and A[r] > A[largest]:
        largest = r

    if largest != i:
        A[i], A[largest] = A[largest], A[i]
        max_heapify(A, heap_size, largest)


def build_max_heap(A):
    """
    building a max heap from sequence A.

    Param: A: a sequence of n numbers.
    """
    heap_size = len(A)
    start = len(A) // 2 - 1
    for i in range(start, -1, -1):
        max_heapify(A, heap_size, i)


def heap_sort(A):
    """
    sort sequence A using heap sort algorthim.
    Param: A: a sequence of n numbers.
    """
    build_max_heap(A)
    heap_size = len(A)
    for i in range(len(A) - 1, 0, -1):
        A[0], A[i] = A[i], A[0]
        heap_size -= 1
        max_heapify(A, heap_size, 0)
