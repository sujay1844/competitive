for _ in range(int(input())):
    n = int(input().strip())
    problems = [int(x) for x in input().strip().split()]
    index = 0
    for i in range(n):
        if 1 <= problems[i] <= 7:
            index += 1
        if index == 7:
            print(i+1)
            break
# took 10 min
