A = [8,2,6,5,1]
n = len(A)
for i in range (n-1):
    position = i
    for j in range (i+1,n):
        if A[j] < A[position]:
           position = j
        A[i],A[position] = A[position],A[i]
        # temp = A[i]
        # A[i] = A[position]
        # A[position] = temp
    print(A)