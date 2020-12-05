#!usr/bin/python3

with open("input") as f:
    data = f.read().splitlines()

pw_valid = 0

for line in data:
    if len(line) == 0:
        break

    x = line.split()

    count = x[0].split('-')
    char = x[1].replace(':', '')
    pwd = x[2]

    c_min = int(count[0])
    c_max = int(count[1])

    if pwd.count(char) >= c_min and pwd.count(char) <= c_max:
        pw_valid = pw_valid + 1

print(pw_valid)
