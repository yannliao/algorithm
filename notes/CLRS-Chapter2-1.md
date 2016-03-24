#Loop invariant:
Loop invariants can help us understand if an algorithm is correct or not. Three thins should be investigated about loop invariant:

* Initialization: It should be true before the first iteration. corresponding to base case of mathmatical induction.
* Maintenance: If it's true before on step then it should be true before the next step. same to the inductive step.
* Termination:: When the loop ends, we should have conditions to determine the algorithm is correct. Different from induction.

#Pseudocod:
Insertion_Sort(A)
    for j = 2 to A.length
        key = A[j]
        i = j-1
        while i > 0 and A[i] > key
          A[i+1] = A[i]
          i = i - 1
    A[i+1] = key

