#!/usr/bin/env python

for i in range(int(input())):
    n = input()
    a = [int(x) for x in input().split(" ")]
    score = 0
    ones = len([x for x in a if x == 1])
    zeroes = len([x for x in a if x == 0])

    if ones >= zeroes:
        score += zeroes
        ones -= zeroes
    else:
        score += ones
        ones = 0

    if ones > 0:
        score += ones // 3
    print(score)


