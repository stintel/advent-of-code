#!usr/bin/python3

with open("input") as f:
    data = f.read().splitlines()

x = 0
trees = 0

for line in data:
    try:
        if line[x] == '#':
            line = line[:x] + 'X' + line[x+1:]
            trees = trees + 1
        else:
            line = line[:x] + 'O' + line[x+1:]

        x = x + 3

        if x >= len(line):
            x = x - len(line)

        print(line)

    except Exception as e:
        print(f"{e}: {x}")
        print(f"line length: {len(line)}")
        break

print(trees)
