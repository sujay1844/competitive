for _ in range(int(input())):
    n, k, v = [int(x) for x in input().split()]

    s = sum([int(x) for x in input().split()])

    x = (v*(n + k) - s)/k
    if (x*10)%10 != 0 or x<=0:
        print(-1)
    else:
        print(int(x))
