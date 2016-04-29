# -*- coding: utf-8 -*-
#
"""
  test code for the maximum subarray problem
"""
import sys
import os
sys.path.insert(0, os.path.abspath('..'))

from python.maximum_subarray import max_subarray


A = [1,4,-2,5,-1,-1,-7,2]
B = max_subarray(A, 0, len(A)-1)
print("the orginal array: ", A)
print("the maximum subarray of array A: ", B)
