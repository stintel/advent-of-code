#!usr/bin/python3

from string import ascii_lowercase


def count_unique_chars(s):
    n = 0

    for c in ascii_lowercase:
        if c in s:
            n = n + 1

    return(n)


with open("input") as f:
    data = f.read().splitlines()

f.close()

count = 0
i = 0
group_answers = ""

for line in data:
    i = i + 1
    group_answers = group_answers + line

    if len(line) == 0 or i == len(data):
        count = count + count_unique_chars(group_answers)
        group_answers = ""

print(count)
