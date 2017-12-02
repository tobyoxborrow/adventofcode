#!/usr/bin/env python3


def get_challenge():
    with open("./input", "r") as pifh:
        line = pifh.readline()
    return line.strip()


def solve(puzzle_input):
    value = 0
    next_digit = None
    length = len(puzzle_input)
    half = int(length / 2)
    for index, digit in enumerate(puzzle_input):
        digit = int(digit)
        next_digit = int(puzzle_input[int(index - half)])
        if digit == next_digit:
            value += digit
    return value


print(solve("1212") == 6)
print(solve("1221") == 0)
print(solve(get_challenge()))
