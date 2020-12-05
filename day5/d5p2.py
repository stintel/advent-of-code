#!usr/bin/python3

def get_col(s):
    col = 0

    for i in range(0, 3):
        if s[i] == 'R':
            col = col | 1 << (2 - i)
        elif s[i] == 'L':
            pass
        else:
            raise Exception("Invalid seat row")

    return(col)

def get_row(s):
    row = 0

    for i in range(0, 7):
        if s[i] == 'B':
            row = row | 1 << (6 - i)
        elif s[i] == 'F':
            pass
        else:
            raise Exception("Invalid seat column")

    return(row)

def get_id(r, c):
    return (r * 8 + c)


with open("input") as f:
    data = f.read().splitlines()

f.close()

seats = [[0 for col in range(8)] for row in range(128)]
maxid = 0

for line in data:
    seat = {}
    if len(line) == 10:
        seats[get_row(line[0:7])][get_col(line[7:10])] = 1
    else:
        raise Exception("Invalid seat")

for row in range(128):
    if seats[row].count(0) == 1:
        col = seats[row].index(0)

        print(f"Seat in col {col}, row {row}, with id {get_id(row, col)} is emtpy: {seats[row]}")
