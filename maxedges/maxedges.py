for _ in range(int(input())):
    n, k, l = [int(x) for x in input().split()]
    srcs = list(range(1, k+1))
    sinks = list(range(n-l+1, n+1))
    print(f"{srcs = }")
    print(f"{sinks = }")
    srcs_filtered = [x for x in srcs if x not in sinks]
    sinks_filtered = [x for x in sinks if x not in srcs]
    print(f"{srcs_filtered = }")
    print(f"{sinks_filtered = }")
