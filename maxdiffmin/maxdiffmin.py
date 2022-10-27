n = int(input().strip())

for _ in range(n):
    str = input().split(" ")
    a = int(str[0])
    b = int(str[1])
    c = int(str[2])
    print(max(a, b, c) - min(a, b, c))
