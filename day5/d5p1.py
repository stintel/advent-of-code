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

seats = []
maxid = 0

for line in data:
    seat = {}
    if len(line) == 10:
        print(line)
        seat['row'] = get_row(line[0:7])
        seat['col'] = get_col(line[7:10])
        seat['id'] = get_id(seat['row'], seat['col'])
        seats.append(seat)
        if seat['id'] > maxid:
            maxid = seat['id']
    else:
        raise Exception("Invalid seat")

print(maxid)
