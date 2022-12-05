A = [2,8,6,5,1]
n = len(A)
for stage in range (n-1,0,-1):
    for i in range (stage):
        if A[i] > A[i+1]:
            A[i],A[i+1] = A[i+1],A[i]
        print(A) 
   

