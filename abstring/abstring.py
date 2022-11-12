from collections import Counter
for _ in range(int(input())):
    input()
    chars = Counter(input())
    for x in chars:
        if chars[x] % 2 == 1:
            print("NO")
            break
    else:
        print("YES")

