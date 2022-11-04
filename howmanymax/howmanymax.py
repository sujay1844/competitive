for _ in range(int(input())):
    n = int(input())
    s = input()
    if len(s) != 1:
        print(s.count("01")+(s.count("10")*2))
    else:
        print(1)
