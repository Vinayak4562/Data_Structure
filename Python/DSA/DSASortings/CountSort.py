def CountSort(arr):
    n = len(arr)
    maxsize = max(arr)
    Carray = [0] *(maxsize+1)
    for i in range (n):
        Carray[arr[i]] = Carray[arr[i]] + 1 
    i=0
    j=0
    while i < maxsize+1:
        if Carray[i] > 0:
            arr[j] = i
            j += 1
            Carray[i] = Carray[i] - 1
        else:
            i += 1
       
arr = [3,5,8,9,6,2,3,5,5]
print('Original Array:', arr)
CountSort(arr)
print('Sorted Array:',arr)