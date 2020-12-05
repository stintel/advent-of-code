#!usr/bin/python3

with open("input") as f:
    data = f.read().splitlines()

for i in range(len(data)):
    for j in range(len(data)):
        for k in range(len(data)):
            sum = int(data[i]) + int(data[j]) + int(data[k])
            if sum == 2020:
                print(f"{data[i]}, {data[j]}, {data[k]}, {int(data[i]) * int(data[j]) * int(data[k])}")
