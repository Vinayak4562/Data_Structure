def BubbleSort(A):
    n = len(A)
    for stage in range (n-1, 0, -1):
        for i in range (stage):
            if A[i] > A[ i+1]:
                temp = A[i]
                A[i] = A[i+1]
                A[i+1] = temp

A = [8,5,1,6,7]
print("Orginal Array: ", A)
BubbleSort(A)
print(" Sorted Array: ", A)