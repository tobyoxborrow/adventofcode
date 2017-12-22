#!/usr/bin/env python3


def get_challenge():
    with open("./input", "r") as pifh:
        line = pifh.readline()
    return line.strip()


def solve(puzzle_input):
    value = 0
    next_digit = None
    for index, digit in enumerate(puzzle_input):
        digit = int(digit)
        try:
            next_digit = int(puzzle_input[index + 1])
        except IndexError:
            next_digit = int(puzzle_input[0])
        if digit == next_digit:
            value += digit
    return value


print(solve("1122") == 3)
print(solve("1111") == 4)
print(solve(get_challenge()))
