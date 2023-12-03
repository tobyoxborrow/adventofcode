#!/usr/bin/env python
# https://adventofcode.com/2023/day/1
#


def get_challenge(filename):
    lines = []
    with open(filename, "r") as fh:
        for line in fh.readlines():
            lines.append(line.strip())
    return lines


def part1(filename):
    lines = get_challenge(filename)

    total = 0
    for line in lines:
        first = 0
        last = 0
        for c in line:
            try:
                first = int(c)
                break
            except ValueError:
                continue
        for c in line[::-1]:
            try:
                last = int(c)
                break
            except ValueError:
                continue
        number = first*10 + last
        total += number
        
    return total


def findDigits(line):
    digits = [
        ('one', 1),
        ('two', 2),
        ('three', 3),
        ('four', 4),
        ('five', 5),
        ('six', 6),
        ('seven', 7),
        ('eight', 8),
        ('nine', 9),
        ('1', 1),
        ('2', 2),
        ('3', 3),
        ('4', 4),
        ('5', 5),
        ('6', 6),
        ('7', 7),
        ('8', 8),
        ('9', 9),
    ]

    positions = {}
    
    for i, _ in enumerate(line):
        for digit in digits:
            try:
                if line[i:].index(digit[0]) == 0:
                    positions[i] = digit[1]
            except ValueError:
                continue
    
    return positions

def part2(filename):
    lines = get_challenge(filename)

    total = 0
    for line in lines:
        positions = findDigits(line)
        first = min(positions.keys())
        last = max(positions.keys())
        number = positions[first]*10 + positions[last]
        total += number
        
    return total


print(part1("sample.txt"))
print(part1("input.txt"))


print(part2("sampleb.txt"))
print(part2("input.txt"))
