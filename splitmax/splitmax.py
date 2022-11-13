for _ in range(int(input())):
    n = int(input())
    s = sum([int(x) for x in input().split()])
    s = s*(s-1)
    print(s % 998244353)
