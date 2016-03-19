"""
  Insertion Sort
  CLRS Chapter 2.1
"""

def sort(sequence):
    """
    Sorting a list of numbers.
    
    Param: sequence: A sequence of n numbers.
    Return: sequence: A permutation(reordering) sequence of the input sequence.
    """
    for j in range(1, len(sequence)):
        key = sequence[j]
        i = j
        while i > 0 and sequence[i - 1] > key:
            sequence[i] = sequence[i - 1]
            i = i - 1
        sequence[i] = key
    return sequence
