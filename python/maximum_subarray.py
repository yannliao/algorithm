"""
  Solve the maximum-subarray problem using divide-and-conquer technique.
  CLRS Chapter 4.1
"""
def max_crossing_subarray(A, low, mid, high):
  left_sum = A[mid]
  left_max = mid
  sum = 0
  for i in range(mid, low-1, -1):
    sum += A[i]
    if sum > left_sum:
      left_sum = sum
      left_max = i
  right_sum = A[mid]
  right_max = mid
  sum = 0
  for j in range(mid, high+1):
    sum += A[j]
    if sum > right_sum:
      right_sum = sum
      right_max = j
  return (left_max, right_max, left_sum+right_sum)

def max_subarray(A, low, high):
  if low == high:
    return (low, high, A[low])
  else:
    mid = int((low + high)/2)
    left_low, left_high, left_sum = max_subarray(A, low, mid)
    right_low, right_high, right_sum = max_subarray(A, mid + 1, high)
    cross_low, cross_high, cross_sum = max_crossing_subarray(A, low, mid, high)
    if (left_sum >= right_sum and left_sum >= cross_sum):
      return (left_low, left_high, left_sum)
    elif (right_sum >= left_sum and right_sum >= cross_sum):
      return (right_low, right_high, right_sum)
    else:
      return (cross_low, cross_high, cross_sum)

