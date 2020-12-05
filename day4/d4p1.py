#!usr/bin/python3


def validate(d):
    #if 'byr' 'iyr' ['eyr'] and d['hgt'] and d['hcl'] and d['ecl'] and d['pid'] and d['cid']:
    if all (k in d for k in ("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")):
        return True

    return False



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
