for i in range(int(input())):
    n = int(input())
    arr = [int(x) for x in input().split(" ")]
    inc = []
    dec = []
    is_inc = False
    for i in range(n):
        for j in range(len(arr[i:])):
            if arr[j] <= arr[j-1] and not is_inc:
                inc.append(arr[j])
            elif arr[j] > arr[j-1]:
                dec.append(arr[j])
                is_inc = True
        print(*inc, *dec)
        inc = []
        dec = []
