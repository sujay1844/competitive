import sys
sys.set_int_max_str_digits(10**7)
facts = {0:1, 1:1}

for _ in range(int(input())):

    input()
    sum = 0

    for x in input().split():
        n = int(x)

        if n not in facts:
            for i in range(1, n+1):
                if i not in facts:
                    facts[i] = facts[i-1] * i

        sum += facts[n]

    sum = sum % (10**9+7)
    print(sum)
