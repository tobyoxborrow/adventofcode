#!/usr/bin/env python3


def main():
    """Main"""
    lines = get_challenge()
    grid = draw_grid(lines)
    print("A: {}".format(solve_a(grid)))
    print("B: {}".format(solve_b(grid)))


def get_challenge():
    """Read lines from the challenge input file"""
    lines = list()
    with open("./input.txt", 'r') as hdl:
        for line in hdl.readlines():
            lines.append(line.strip())
    return lines


def parse_claim(claim):
    """
    Parse a line of input

    Example claim:
    #1 @ 861,330: 20x10
    """

    tokens = claim.split(" ")
    claim_id = int(tokens[0][1:])
    (offset_x, offset_y) = tokens[2].split(',')
    offset_x = int(offset_x)
    offset_y = int(offset_y[:-1])
    (width, height) = tokens[3].split('x')
    width = int(width)
    height = int(height)

    return (claim_id, offset_x, offset_y, width, height)


def draw_grid(claims):
    """Draw the grid and the claims"""
    # create a dict of tuple coordinates
    # seems a bit easier than a list of lists
    grid = dict()
    for grid_x in range(1000):
        for grid_y in range(1000):
            grid[(grid_x, grid_y)] = list()

    for claim in claims:
        parsed_claim = parse_claim(claim)
        draw_claim(grid, parsed_claim)

    return grid


def draw_claim(grid, claim):
    """
    Draw a claim on the grid

    If the inch is claimed, store a list of ids that claim it
    """

    (claim_id, offset_x, offset_y, width, height) = claim

    for claim_x in range(width):
        for claim_y in range(height):
            grid[(offset_x+claim_x, offset_y+claim_y)].append(claim_id)


def solve_a(grid):
    """Solve Part A"""
    overlapping = 0

    for pos in grid:
        # count how many ids are stored at each position
        # any inches with more than one claim id are considered overlapping
        if len(grid[pos]) > 1:
            overlapping += 1

    return overlapping


def solve_b(grid):
    """Solve Part B"""
    seen_claims = set()
    overlapping_claims = set()

    for pos in grid:
        # go through the grid and record each claim seen
        # and record each overlapping claim
        claim_ids = grid[pos]
        if not claim_ids:
            continue
        for claim_id in claim_ids:
            seen_claims.add(claim_id)
        if len(claim_ids) > 1:
            for claim_id in claim_ids:
                overlapping_claims.add(claim_id)

    for overlapping_claim_id in overlapping_claims:
        seen_claims.remove(overlapping_claim_id)

    if len(seen_claims) != 1:
        raise Exception("Found too many non-overlapping claims")

    return seen_claims.pop()


if __name__ == '__main__':
    main()
