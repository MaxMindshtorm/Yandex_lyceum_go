n = 10**8 - 172
i = 1
while n > 0:
    n = n - 2 ** (3 * i - 1)
    print(n)
    i += 1
print(i)