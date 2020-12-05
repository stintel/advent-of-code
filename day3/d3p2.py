#!usr/bin/python3

def traverse(data, r,d):
    x = 0
    y = 0
    trees = 0

    while y < len(data):
        if data[y][x] == '#':
            #data[y] = data[y][:x] + 'X' + data[y][x+1:]
            trees = trees + 1
        #else:
            #data[y] = data[y][:x] + 'O' + data[y][x+1:]

        #print(data[y])

        x = x + r

        if x >= len(data[y]):
            x = x - len(data[y])

        y = y + d

    print(f"Trees: {trees}")
    return(trees)


if __name__ == '__main__':
    with open("input") as f:
        content = f.read().splitlines()

    f.close()

    t1 = traverse(content, 1, 1)
    t2 = traverse(content, 3, 1)
    t3 = traverse(content, 5, 1)
    t4 = traverse(content, 7, 1)
    t5 = traverse(content, 1, 2)

    print(t1 * t2 * t3 * t4 * t5)
