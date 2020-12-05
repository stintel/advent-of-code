#!usr/bin/python3

def validate(d):
    try:
        byr = d['byr']
        iyr = d['iyr']
        eyr = d['eyr']
        hgt = d['hgt']
        hcl = d['hcl']
        ecl = d['ecl']
        pid = d['pid']
    except Exception as e:
        return False

    if not (len(byr) == 4 and int(byr) >= 1920 and int(byr) <= 2002):
        return False

    if not (len(iyr) == 4 and int(iyr) >= 2010 and int(iyr) <= 2020):
        return False

    if not (len(eyr) == 4 and int(eyr) >= 2020 and int(eyr) <= 2030):
        return False

    if hgt.endswith('cm'):
        cm = int(hgt.rstrip('cm'))
        if not (cm >= 150 and cm <= 193):
            return False
    elif hgt.endswith('in'):
        inch = int(hgt.rstrip('in'))
        if not (inch >= 59 and inch <= 76):
            return False
    else:
        return False

    if hcl.startswith('#'):
        hexval = hcl.lstrip('#')
        if not len(hexval) == 6:
            return False
        try:
            int(hexval, 16)
        except Exception as e:
            print(e)
            return False
    else:
        return False

    if not ecl in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
        return False

    if len(pid) == 9:
        try:
            int(pid)
        except Exception as e:
            print(e)
            return False
    else:
        return False

    return True


with open("input") as f:
    data = f.read().splitlines()

f.close()

passport_strings = []

i = 0
s = ''

for line in data:
    s = s + line + " "
    i = i + 1
    
    if len(line) == 0 or i == len(data):
        passport_strings.append(s)
        s = ''

passports = []

for ps in passport_strings:
    p = {}
    fields = ps.split()

    for f in fields:
        t = f.split(':')
        k = t[0]
        v = t[1]

        p[k] = v

    passports.append(p)

valid = 0

for p in passports:
    if validate(p):
        valid = valid + 1

print(valid)
