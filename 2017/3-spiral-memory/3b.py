#!/usr/bin/env python3
"""

Day 3: Spiral Memory

147  142  133  122   59
304    5    4    2   57
330   10    1    1   54
351   11   23   25   26
362  747  806--->   ...

What is the first value written that is larger than your puzzle input?

"""


def get_challenge():
    """Get the challenge input"""
    with open("./input", "r") as pifh:
        line = pifh.readline()
    return int(line.strip())


def solve(puzzle_input):
    """Solve a puzzle input"""
    # print("solve(%s)" % puzzle_input)
    direction = 0       # 0 up, 1 left, 2 down, 3 right

    # initialise a starting grid with the first ring aka square 1
    # The size is known to be big enough to solve my input
    grid_size = 13
    grid = [[0 for x in range(grid_size)] for y in range(grid_size)]
    grid[int(grid_size/2)][int(grid_size/2)] = 1
    csx = int(grid_size/2)+1  # current square x co-ordinates within the grid
    csy = int(grid_size/2)    # current square y co-ordinates within the grid

    while True:
        # set the square value based on our neighboughs
        grid[csy][csx] = (
            grid[csy-1][csx-1] +
            grid[csy-1][csx] +
            grid[csy-1][csx+1] +
            grid[csy][csx-1] +
            grid[csy][csx] +
            grid[csy][csx+1] +
            grid[csy+1][csx-1] +
            grid[csy+1][csx] +
            grid[csy+1][csx+1]
            )
        if grid[csy][csx] > puzzle_input:
            break

        # do we need to change direction?
        # if we are going up and there is nothing left, turn left
        if direction == 0:
            if not grid[csy][csx-1]:
                direction = 1
        # if we are going left and there is nothing below, turn down
        elif direction == 1:
            if not grid[csy+1][csx]:
                direction = 2
        # if we are going down and there is nothing right, turn right
        elif direction == 2:
            if not grid[csy][csx+1]:
                direction = 3
        # if we are going right and there is nothing above, turn up
        elif direction == 3:
            if not grid[csy-1][csx]:
                direction = 0

        # move to the next square
        if direction == 0:
            csy -= 1
        elif direction == 1:
            csx -= 1
        elif direction == 2:
            csy += 1
        elif direction == 3:
            csx += 1

    # for row in grid:
    #     print(row)
    return grid[csy][csx]


print(solve(1) == 2)
print(solve(3) == 4)
print(solve(4) == 5)
print(solve(5) == 10)
print(solve(15) == 23)
print(solve(get_challenge()))
