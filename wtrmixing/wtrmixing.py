def check(a:int, b:int, x:int, y:int):
    if a < b and b-a <= x: return True
    if a > b and a-b <= y: return True
    if a == b: return True
    return False

for _ in range(int(input().strip())):
    a, b, x, y = [int(x) for x in input().strip().split()]
    print("YES" if check(a, b, x, y) else "NO")
