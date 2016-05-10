# -*- coding: utf-8 -*-
"""
  sample code for the heap sort
"""

import sys
import os
sys.path.insert(0, os.path.abspath('..'))

from python.quick_sort import quicksort, randomized_quicksort

A = [2,8,7,1,3,5,6,4]
print("the orginal array: ", A)
quicksort(A, 0, len(A) - 1)
print("after quick sort: ", A)
B = [2,8,7,1,3,5,6,4]
randomized_quicksort(B, 0, len(B) - 1)
print("after randomized quick sort: ", B)
