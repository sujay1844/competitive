for _ in range(int(input().strip())):
    n = int(input().strip())
    arr = [int(x) for x in input().strip().split()]
    count  = 0
    for i in range(1, n+1):
        if arr[i-1] == (i + count):
            count += 1
    print(count)
