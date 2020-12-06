#!usr/bin/python3

from string import ascii_lowercase


def init_dict():
    answer = {}

    for c in ascii_lowercase:
        answer[c] = 0

    return(answer)


with open("input") as f:
    data = f.read().splitlines()

f.close()


answer = init_dict()
count = 0
j = 0

for line in data:
    print(line)
    if len(line) == 0:
        for c in ascii_lowercase:
            if answer[c] == j:
                count = count + 1
        answer = init_dict()
        j = 0
    else:
        j = j + 1
        for c in line:
            answer[c] = answer[c] + 1
            print(f"{c}:{answer[c]}")

for c in ascii_lowercase:
    if answer[c] == j:
        count = count + 1

print(count)
