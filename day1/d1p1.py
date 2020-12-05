#!usr/bin/python3

with open("input") as f:
    data = f.read().splitlines()

for i in range(len(data)):
    for j in range(len(data)):
        sum = int(data[i]) + int(data[j])
        if sum == 2020:
            print(f"{data[i]}, {data[j]}, {int(data[i]) * int(data[j])}")
