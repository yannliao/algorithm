# -*- coding: utf-8 -*-
"""
  sample code for the count sort
"""

import context
from python.count_sort import counting_sort

array = [1, 4, 2, 5, 0, 10, 7, 2]
bucket = []
print("the orginal array: ", array)
counting_sort(array, bucket, 11)
print("after: ", bucket)
