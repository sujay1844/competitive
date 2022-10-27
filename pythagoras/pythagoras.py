t = int(input())
for _ in range(t):
    flag = True
    n = int(input())
    for i in range(int(n/2)):
        for j in range(i, int(n/2)):
            if (i*i + j*j) == n:
                print(i, j)
                flag = False
    if not flag:
        print(-1)
