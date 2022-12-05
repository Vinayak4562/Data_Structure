def RadixSort(arr):
    n = len(arr)
    maxEle = max(arr)
    digits = len(str(maxEle))
    l = []
    bins = [l] * 10
    for i in range (digits):
        for j in range (n):
            e = int((arr[j] / pow(10,i)) % 10)
            if len(bins[e]) > 0:
                bins[e].append(arr[j])
            else:
                bins[e] = [arr[j]]
        k = 0

        for x in range (10):
            if len(bins[x]) > 0:
                for y in range(len(bins[x])):
                    arr[k] = bins[x].pop(0)
                    k += 1

arr = [244,682,73,6,123, 820,200,10,2]
print("Original arry", arr)
RadixSort(arr)
print("Sorted arry", arr)