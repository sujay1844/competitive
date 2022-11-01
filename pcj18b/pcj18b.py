import math
for _ in range(int(input().strip())):
    n = int(input().strip())
    ans = 0
    for i in range(1, n+1, 2):
        ans += math.factorial(n-i+1)
            
    # ans  = ((n**3)/2.0) + (n**2.0) + (n/2.0) - (1.0/3.0)*(n)*(n+1)*(n+2)
    print(int(ans))
