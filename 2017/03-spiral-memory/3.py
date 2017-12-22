#!/usr/bin/env python3


def get_challenge():
    with open("./input", "r") as pifh:
        line = pifh.readline()
    return int(line.strip())


def solve(puzzle_input):
    print("solve(%s)" % puzzle_input)
    if puzzle_input == 1:
        return 0
    c = 1           # ring number
    previous = 1    # last number in the previous ring
    while True:
        squares_in_ring = (c * 8)
        ring_end_square = previous + squares_in_ring
        # are we on a ring that contains the puzzle input?
        if ring_end_square > puzzle_input:
            # now we know how many rings away from the access square we are (c)
            # we next need to work out the offset from the input to one of the
            # north/south/east/west pathways
            south_square = ring_end_square - c
            west_square = south_square - (squares_in_ring / 4)
            north_square = west_square - (squares_in_ring / 4)
            east_square = north_square - (squares_in_ring / 4)
            print(squares_in_ring)
            print(south_square, west_square, north_square, east_square)
            offsets = [
                abs(puzzle_input - south_square),
                abs(puzzle_input - west_square),
                abs(puzzle_input - north_square),
                abs(puzzle_input - east_square),
                ]
            offset = min(offsets)
            print(c, offset)
            return int(c + offset)
        previous = ring_end_square
        c += 1


print(solve(1) == 0)
print(solve(12) == 3)
print(solve(23) == 2)
print(solve(1024) == 31)
print(solve(get_challenge()))
