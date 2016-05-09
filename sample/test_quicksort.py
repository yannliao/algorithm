# -*- coding: utf-8 -*-
"""
  sample code for the heap sort
"""

import sys
import os
sys.path.insert(0, os.path.abspath('..'))

from python.quick_sort import quicksort

A = [2,8,7,1,3,5,6,4]
print("the orginal array: ", A)
quicksort(A, 0, len(A) - 1)
print("after: ", A)
