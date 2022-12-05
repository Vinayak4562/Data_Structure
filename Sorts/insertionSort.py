def InertionSort (A):
    n = len(A)
    for i in range (1, n):
        Cvalue = A[i]
        position = i
        while position > 0 and A[position-1] > Cvalue:
            A[position] = A[position-1]
            position = position-1
        A[position] = Cvalue


A = [3,5,8,9,6,2]
print("Original Array: ", A)
InertionSort(A)
print("Sorted Array:", A)
