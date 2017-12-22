#!/usr/bin/env python3


def get_challenge():
    with open("./input", "r") as input_fh:
        return input_fh.readlines()


def solve(puzzle_input):
    checksum = 0
    for line in puzzle_input:
        line = line.strip()
        if not line:
            continue
        tokens = line.split("\t")
        for token in tokens:
            digit = int(token)
            for token2 in tokens:
                digit2 = int(token2)
                if digit == digit2:
                    continue
                result = digit % digit2
                if result == 0:
                    checksum += digit / digit2
                    break
    return int(checksum)


print(solve(
    [
        "5	9	2	8",
        "9	4	7	3",
        "3	8	6	5"
        ]
    ) == 9)
print(solve(get_challenge()))
