#!/usr/bin/env python3
"""

Day 4: High-Entropy Passphrases

abcde fghij is a valid passphrase.
abcde xyz ecdab is not valid - the letters from the third word can be
rearranged to form the first word.  a ab abc abd abf abj is a valid passphrase,
because all letters need to be used when forming another word.
iiii oiii ooii oooi oooo is valid.
oiii ioii iioi iiio is not valid - any of these words can be rearranged to form
any other word.

Under this new system policy, how many passphrases are valid?

"""


def get_challenge():
    with open("./input", "r") as input_fh:
        return input_fh.readlines()


def solve(puzzle_input):
    valid_count = 0
    for line in puzzle_input:
        line = line.strip()
        if not line:
            continue
        tokens = line.split(" ")
        words = set()
        is_valid = True
        for token in tokens:
            sorted_token = ''.join(sorted([c for c in token]))
            if sorted_token in words:
                is_valid = False
                break
            words.add(sorted_token)
        if is_valid:
            valid_count += 1
    return valid_count


print(solve(
    [
        "abcde fghij",
        ]
    ) == 1)
print(solve(
    [
        "abcde fghij",
        "abcde xyz ecdab",
        "a ab abc abd abf abj",
        ]
    ) == 2)
print(solve(get_challenge()))
