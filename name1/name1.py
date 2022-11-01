from collections import Counter
for _ in range(int(input())):
    parent = input().replace(" ", "")
    parent_parsed = Counter(parent)
    number_of_children = int(input())
    children = ""
    for _ in range(number_of_children):
        children += input()
    children_parsed = Counter(children)
    flag = True
    for i in children:
        if i in parent:
            if children_parsed[i] > parent_parsed[i]:
                flag = False
        else:
            flag = False
            break
    print("YES" if flag else "NO")
