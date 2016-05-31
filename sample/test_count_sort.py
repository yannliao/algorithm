# -*- coding: utf-8 -*-
"""
  sample code for the count sort
"""

import context
from python.count_sort import counting_sort

array = [2, 5, 3, 0, 2, 3, 0, 3]
print("the orginal array: ", array)
bucket = counting_sort(array,  6)
print("after: ", list(bucket.values()))
