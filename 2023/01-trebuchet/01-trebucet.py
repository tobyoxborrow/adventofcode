#!/usr/bin/env python
# https://adventofcode.com/2023/day/1
#


def get_challenge(filename):
    lines = []
    with open(filename, "r") as fh:
        for line in fh.readlines():
            lines.append(line.strip())
    return lines


def solve(filename):
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


print(solve("sample.txt"))
print(solve("input.txt"))
