#!/usr/bin/python3
"""
  test code for the maximum subarray problem
"""
from maximum_subarray import *

A = [1,4,-2,5,-1,-1,-7,2]
B = max_subarray(A, 0, len(A)-1)
print("the orginal array: ", A)
print("the maximum subarray of array A: ", B)
