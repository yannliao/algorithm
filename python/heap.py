# -*- coding: utf-8 -*-
"""
  Heap Sort
  CLRS Chapter 6
"""

def max_heapify(A, heap_size, i):
  """
  lets the value A[i] "float down" in the max-heap, so that the subtree at index i obeys the max-heap property.

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
  heap_size = len(A)
  start = len(A) // 2 - 1
  for i in range(start, -1, -1):
    max_heapify(A, heap_size, i)


def heap_sort(A):
  build_max_heap(A)
  heap_size = len(A)
  for i in range(len(A) - 1, 0, -1):
    A[0], A[i] = A[i], A[0]
    heap_size -= 1
    max_heapify(A, heap_size, 0)




