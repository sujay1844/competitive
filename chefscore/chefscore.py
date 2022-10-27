t = int(input())
for _ in range(t):
    # n, x, y = [int(a) for a in input().split()]
    n, x, y = map(int, input().split())
    if y % x == 0 and n*x >= y:
        print("YES")
    else:
        print("NO")
