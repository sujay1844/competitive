t = int(input())
for _ in range(t):
    n, q = map(int, input().split())
    arr_sum = sum([int(x) for x in input().split()])
    for _ in range(q):
        l, r = [int(x) for x in input().split()]
        if (r-l+1) % 2 != 0:
            arr_sum += 1
    print(arr_sum)
