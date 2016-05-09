#some algorithms

# Chapter 2
## Loop invariant:
Loop invariants can help us understand if an algorithm is correct or not. Three thins should be investigated about loop invariant:

* Initialization: It should be true before the first iteration. corresponding to base case of mathmatical induction.
* Maintenance: If it's true before on step then it should be true before the next step. same to the inductive step.
* Termination:: When the loop ends, we should have conditions to determine the algorithm is correct. Different from induction.

## Pseudocodof Insertion Sort
    Insertion_Sort(A)
        for j = 2 to A.length
            key = A[j]
            i = j-1
            while i > 0 and A[i] > key
              A[i+1] = A[i]
              i = i - 1
        A[i+1] = key
        
## Divide and Conquer
Divide the complicate problem to many small and fimilar problems, sovle subproblems recursively, and combine result of subproblems to get the result of the original problem.

* Divide: divide the problem into many subproblems that are smaller instances of the same problem.
* Conquer: solve subproblems recursively , mostly useing same method. If the subproblem is small enough then solve it directly. 
* Combine: combine the reuslt of subproblems to get the result of the original one.
