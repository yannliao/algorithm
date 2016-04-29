# -*- coding: utf-8 -*-
"""
  Priority queues
  CLRS Chapter 6
"""
from heap import max_heapify

def heap_maximum(A):
  return A[0]

def heap_extract_max(A):
  heap_size = len(A)
  if heap_size < 1:
    print("heap underflow")
    return
  max = A[0]
  A[0] = A[heap_size]
  heap_size -= 1
  max_heapify(A, heap_size, 0)
  return max
