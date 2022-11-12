import math

facts = {}

for _ in range(int(input())):

    input()
    sum = 0

    for x in input().split():
        n = int(x)

        try:
            sum += facts[n]

        except KeyError:

            for i in range(n, 0, -1):
                if i in facts: break
                else: facts[i] = math.factorial(i)

            sum += facts[n]

    print(sum % 1_000_000_007)
