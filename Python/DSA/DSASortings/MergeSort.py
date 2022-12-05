def mergesort(A, left, right):
    if left < right:
        mid = (left + right) // 2
        mergesort(A, left, mid)
        mergesort(A, mid+1, right)
        merge(A, left, mid, right)

def merge(A, left, mid, right):
    i = left
    j = mid+1
    k = left
    B = [0] * (right+1)
    # print("after mearging ",A)
    # print()
    while i <= mid and j <= right:
        if A[i] < A[j]:
            B[k] = A[i]
            i = i + 1
        else:
            B[k] = A[j]
            j = j + 1
        k = k + 1

    while i <= mid:
        B[k] = A[i]
        i = i + 1
        k = k + 1
    # print("After merging left remaining Ele",A)
    # print()

    while j <= right:
        B[k] = A[j]
        j = j + 1
        k = k + 1
    for x in range(left,right+1):
        A[x] = B[x]
    # print("After merging right remaining Ele",A)
    # print()


A = [2,5,9,7,4,6,2,5,8,2]
print('Original Array:',A)
mergesort(A,0,len(A)-1)
print('Sorted Array:',A)
