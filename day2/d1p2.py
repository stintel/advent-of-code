#!usr/bin/python3

with open("input") as f:
    data = f.read().splitlines()

pw_valid = 0

for line in data:
    if len(line) == 0:
        break

    x = line.split()

    pos = x[0].split('-')
    char = x[1].replace(':', '')
    pwd = x[2]

    pos_val = int(pos[0]) - 1
    pos_inv = int(pos[1]) - 1

    print(line)
    print(f"{char}, {pos_val}, {pos_inv}, {pwd}")
    print(f"{pwd[pos_val]}, {pwd[pos_inv]}")

    val = 0
    if pwd[pos_val] == char:
        val = val + 1

    if pwd[pos_inv] == char:
        val = val + 1

    if val == 1:
        pw_valid = pw_valid + 1
        print("Valid")

    print()

print(pw_valid)
