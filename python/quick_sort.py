# -*- coding: utf-8 -*-
"""
  Quick Sort
  CLRS Chapter 7.1
"""

def partition(A, p, r):
  """
    partitioning the array.
    Param A: arry or subarray to be partition.
    Param p: low index.
    Param r: high index.
  """
  x = A[r] #pivot element
  i = p - 1
  for j in range(p, r, 1):
    if A[j] <= x:
      i = i + 1
      A[i], A[j] = A[j], A[i]
  A[i + 1], A[r] = A[r], A[i + 1]
  print('A={0}, x={1}, i+1={2}, p={3}, r={4} '.format(A, x,  i + 1, p, r))
  return i + 1


def quicksort(A, p, r):
  """
    sort array A using Quicksort.
    Param A: array to be sorted.
    Param p: low index.
    Param r: high index.
  """
  if p < r:
    q = partition(A, p, r)
    print('q={0}'.format(q))
    quicksort(A, p, q - 1)
    quicksort(A, q + 1, r)





