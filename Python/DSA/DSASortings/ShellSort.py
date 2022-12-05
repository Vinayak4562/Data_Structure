def ShellSort(arr):
    n = len(arr)
    for gap in range (n/2):
        for i in range (gap):
            gvalue = arr[i]
            j = i - gap
            while j >= 0 and arr[j] > gvalue:
                arr[j+gap] = arr[j]
                j = j - gap
            arr[j+gap] = gvalue

arr = [3,5,8,9,6,2]
print(arr)
print(ShellSort(arr))
print(arr)
