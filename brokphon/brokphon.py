def main():
    for _ in range(int(input())):
        n = int(input())
        numbers = [int(x) for x in input().split()]
        count = 0
        for i in range(1, n-1):
            if numbers[i] != numbers[i-1] or numbers[i] != numbers[i+1]:
                count += 1
        if numbers[0] != numbers[1]:
            count += 1
        if numbers[n-1] != numbers[n-2]:
            count += 1
        print(count)
    

if __name__ == '__main__':
    main()
