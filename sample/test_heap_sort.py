# -*- coding: utf-8 -*-
"""
  sample code for the heap sort
"""

import context
from python.heap import heap_sort


A = [1, 4, 2, 5, 0, 10, 7, 2]
print("the orginal array: ", A)
heap_sort(A)
print("after: ", A)
